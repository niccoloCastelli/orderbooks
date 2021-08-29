package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/config"
	"github.com/niccoloCastelli/orderbooks/server/proto"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/niccoloCastelli/orderbooks/workers"
	"github.com/niccoloCastelli/orderbooks/workers/backends/csv_backend"
	"github.com/niccoloCastelli/orderbooks/workers/downloaders"
	"github.com/niccoloCastelli/orderbooks/workers/writers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"time"
)

const (
	eventsBasePath    = "events"
	eventsFileMatchRe = `^(\w+)\/(.*?)\/(\d{4})\/(\d{2})\/(\d{2})\/([A-Z]+-[A-Z]+)\.csv\.gz$`
)

type orderbooksGrpcServer struct {
	Logger      *zerolog.Logger
	Storage     afero.Fs
	Config      config.ServerConfig
	exchanges   []*orderbooks.ExchangeMsg
	reader      workers.SnapshotReader
	cacheWriter *writers.CacheWriter
}

func newOrderbooksGrpcServer(logger *zerolog.Logger, conf config.ServerConfig) (*orderbooksGrpcServer, error) {
	if logger == nil {
		logger = &log.Logger
	}

	if _, err := os.Stat(conf.StoragePath); os.IsNotExist(err) {
		log.Error().Err(err).Send()
		return nil, err
	}
	fs := afero.NewBasePathFs(afero.NewOsFs(), conf.StoragePath)
	reader, err := csv_backend.NewSnapshotReader(fs, logger)
	if err != nil {
		return nil, err
	}

	pair := common.Pair{Base: "BTC", Quote: "USD"}
	snapshotChan := make(chan common.Snapshot, 4)
	cacheWriter := writers.NewCacheWriterWriter(*logger, "bitmex", common.Pair{Base: "BTC", Quote: "USD"})
	sd := downloaders.NewSnapshotDownloader([]common.Pair{pair}, time.Second*60, snapshotChan, 10)
	if err := sd.Init("bitmex"); err != nil {
		return nil, err
	}
	if err := sd.Run(context.Background()); err != nil {
		return nil, err
	}
	cacheWriter.Init(snapshotChan)
	cacheWriter.Run(context.Background())

	return &orderbooksGrpcServer{
		Logger:      logger,
		Config:      conf,
		Storage:     fs,
		reader:      reader,
		cacheWriter: cacheWriter,
	}, nil
}
func (w orderbooksGrpcServer) GetExchanges(ctx context.Context, msg *orderbooks.EmptyMsg) (*orderbooks.GetExchangesResponseMsg, error) {
	exchangesList := w.reader.Exchanges()
	exchanges := make([]*orderbooks.ExchangeMsg, len(exchangesList))
	for i, exchangeName := range exchangesList {
		pairs := w.reader.Pairs(exchangeName)
		exchanges[i] = &orderbooks.ExchangeMsg{
			Id:    exchangeName,
			Name:  exchangeName,
			Pairs: make([]*orderbooks.PairMsg, len(pairs)),
		}
		for j, pair := range pairs {
			dateRange := w.reader.TimeRange(exchangeName, pair)
			if dateRange == nil {
				w.Logger.Error().Str("exchange", exchangeName).Str("pair", pair.String()).Err(errors.New("empty range")).Send()
				continue
			}
			exchanges[i].Pairs[j] = &orderbooks.PairMsg{
				Base:      pair.Base,
				Quote:     pair.Quote,
				DateStart: utils.TimeToGogoProtoTsPtr(&dateRange.Start),
				DateEnd:   utils.TimeToGogoProtoTsPtr(&dateRange.End),
			}
		}
	}
	w.exchanges = exchanges
	return &orderbooks.GetExchangesResponseMsg{
		Count:     uint32(len(exchanges)),
		Exchanges: exchanges,
	}, nil
}

func (w orderbooksGrpcServer) QueryEvents(msg *orderbooks.EventsQueryMsg, server orderbooks.OrderBooks_QueryEventsServer) error {
	pair, err := common.PairsFromString(msg.Pair)
	if err != nil {
		return err
	}
	tr := w.reader.TimeRange(msg.Exchange, *pair)
	if tr == nil {
		return status.New(codes.NotFound, "exchange or pair not found").Err()
	}
	dateStart := utils.GogoProtoTsToTimePtr(msg.DateStart)
	if dateStart == nil {
		dateStart = &tr.Start
	}
	dateEnd := utils.GogoProtoTsToTimePtr(msg.DateEnd)
	if dateEnd == nil {
		dateEnd = &tr.End
	}
	if dateStart.Before(tr.Start) || dateEnd.After(tr.End) {
		w.Logger.Error().Str("time_start", dateStart.String()).Str("time_end", dateEnd.String()).Msg("invalid time range")
		return status.New(codes.OutOfRange, "invalid time range").Err()
	}
	logger := w.Logger.With().Str("exchange", msg.Exchange).Str("pair", pair.String()).Time("start", *dateStart).Time("end", *dateEnd).Logger()
	readRange := utils.NewTimeRange(*dateStart, *dateEnd)
	logger.Debug().Msg("start reading")
	var (
		interval     int64
		snapshotMode common.SnapshotMode
	)
	switch v := msg.SnapshotInterval.(type) {
	case *orderbooks.EventsQueryMsg_Interval:
		intervalD, err := time.ParseDuration(v.Interval)
		if err != nil {
			return err
		}
		interval = int64(intervalD)
		snapshotMode = common.SnapshotModeTime
	case *orderbooks.EventsQueryMsg_Ticks:
		interval = v.Ticks
		snapshotMode = common.SnapshotModeTicks
	}
	ch, err := w.reader.Read(msg.Exchange, *pair, readRange, interval, snapshotMode, int(msg.SnapshotSize), server.Context())
	if err != nil {
		return err
	}
	for {
		select {
		case snapshot, ok := <-ch:
			if !ok {
				logger.Debug().Msg("channel closed")
				return nil
			}
			pbOrders := make([]*orderbooks.Event, len(snapshot.Orders))
			for i, order := range snapshot.Orders {
				pbOrders[i] = &orderbooks.Event{
					Timestamp: order.Timestamp.UnixNano(),
					OrderSide: orderbooks.OrderSide(order.Side),
					Amount:    float32(order.Amount),
					Price:     float32(order.Price),
				}
			}
			if err := server.Send(&orderbooks.SnapshotMsg{
				Timestamp: utils.TimeToGogoProtoTsPtr(&snapshot.Timestamp),
				Exchange:  snapshot.Exchange,
				Pair:      snapshot.Pair.String(),
				Orders:    pbOrders,
				Events:    nil,
			}); err != nil {
				logger.Err(err).Send()
				return nil
			}
		case <-server.Context().Done():
			logger.Debug().Msg("server closed")
			return io.EOF
		}
	}
}
func (w orderbooksGrpcServer) GetLiveData(msg *orderbooks.EventsQueryMsg, server orderbooks.OrderBooks_GetLiveDataServer) error {
	pair, err := common.PairsFromString(msg.Pair)
	if err != nil {
		return err
	}
	timeStart := time.Now()
	logger := w.Logger.With().Str("exchange", msg.Exchange).Str("pair", pair.String()).Time("start", timeStart).Logger()
	logger.Debug().Msg("start live data feed")
	var (
		interval     int64
		snapshotChan = make(chan common.Snapshot, 1024)
		sd           downloaders.Downloader
	)
	switch v := msg.SnapshotInterval.(type) {
	case *orderbooks.EventsQueryMsg_Interval:
		intervalD, err := time.ParseDuration(v.Interval)
		if err != nil {
			return err
		}
		interval = int64(intervalD)
		sd = downloaders.NewSnapshotDownloader([]common.Pair{*pair}, time.Duration(interval), snapshotChan, int(msg.SnapshotSize))
	case *orderbooks.EventsQueryMsg_Ticks:
		interval = v.Ticks
		sd = downloaders.NewSnapshotDownloaderTicks([]common.Pair{*pair}, int(interval), snapshotChan, int(msg.SnapshotSize))
	}

	if err := sd.Init(msg.Exchange); err != nil {
		return err
	}
	if err := sd.Run(server.Context()); err != nil {
		return err
	}
	for {
		select {
		case snapshot, ok := <-snapshotChan:
			if !ok {
				logger.Debug().Msg("channel closed")
				return nil
			}
			pbOrders := make([]*orderbooks.Event, len(snapshot.Orders))
			for i, order := range snapshot.Orders {
				pbOrders[i] = &orderbooks.Event{
					Timestamp: order.Timestamp.UnixNano(),
					OrderSide: orderbooks.OrderSide(order.Side),
					Amount:    float32(order.Amount),
					Price:     float32(order.Price),
				}
			}
			if err := server.Send(&orderbooks.SnapshotMsg{
				Timestamp: utils.TimeToGogoProtoTsPtr(&snapshot.Timestamp),
				Exchange:  snapshot.Exchange,
				Pair:      snapshot.Pair.String(),
				Orders:    pbOrders,
				Events:    nil,
			}); err != nil {
				logger.Err(err).Send()
				return nil
			}
		case <-server.Context().Done():
			logger.Debug().Msg("server closed")
			return io.EOF
		}
	}
}
func (w orderbooksGrpcServer) GetCachedData(msg *orderbooks.EventsQueryMsg, server orderbooks.OrderBooks_GetCachedDataServer) error {
	chanId := fmt.Sprintf("%p", server.Context())
	w.Logger.Info().Str("chan_id", chanId).Msg("cached data stream created")
	ch := w.cacheWriter.SubscribeUpdates(chanId, 4)
	defer w.cacheWriter.UnsubscribeUpdates(chanId)
	currentSnapshots := w.cacheWriter.GetSnapshots()
	for _, snapshot := range currentSnapshots {
		if err := server.Send(SnapshotToPb(&snapshot)); err != nil {
			return err
		}
	}
	for {
		select {
		case snapshot, ok := <-ch:
			if !ok {
				return io.EOF
			}
			if err := server.Send(SnapshotToPb(snapshot)); err != nil {
				w.Logger.Err(err).Send()
				return err
			}
			continue
		case <-server.Context().Done():
			return nil
		}
	}
}
