package downloaders

import (
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/exchanges"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/niccoloCastelli/orderbooks/exchanges/bitmex"
	_ "github.com/niccoloCastelli/orderbooks/exchanges/bitstamp"
	_ "github.com/niccoloCastelli/orderbooks/exchanges/coinbase"
	_ "github.com/niccoloCastelli/orderbooks/exchanges/kraken"
	_ "github.com/niccoloCastelli/orderbooks/exchanges/okex"

	"context"
	"errors"
	"fmt"
	"github.com/niccoloCastelli/orderbooks/utils"
	"os"
	"strings"
	"time"
)

const (
	DefaultSnapshotInterval = time.Second * 30
	maxRetryCount           = 20
	retryInterval           = time.Second * 5
)

func NewSnapshotDownloader(pairs []common.Pair, snapshotInterval time.Duration, outChan chan common.Snapshot, maxSnapshotSize int) Downloader {
	fmt.Println("SnapshotDownloader pairs: ", pairs)
	return &SnapshotDownloader{snapshotInterval: int(snapshotInterval), pairs: pairs, outChan: outChan, snapshotMode: common.SnapshotModeTime, maxSnapshotSize: maxSnapshotSize, cfg: SnapshotDownloaderConfig{
		ExitOnClose: true,
	}}
}
func NewSnapshotDownloaderTicks(pairs []common.Pair, snapshotInterval int, outChan chan common.Snapshot, maxSnapshotSize int) Downloader {
	fmt.Println("SnapshotDownloader pairs: ", pairs)
	return &SnapshotDownloader{snapshotInterval: snapshotInterval, pairs: pairs, outChan: outChan, snapshotMode: common.SnapshotModeTicks, maxSnapshotSize: maxSnapshotSize, cfg: SnapshotDownloaderConfig{
		ExitOnClose: false,
	}}
}

type SnapshotDownloaderConfig struct {
	ExitOnClose bool
}

type SnapshotDownloader struct {
	exchanges        []common.Exchange
	snapshotInterval int
	pairs            []common.Pair
	ctx              context.Context
	cancelFunc       context.CancelFunc
	outChan          chan common.Snapshot
	snapshotMode     common.SnapshotMode
	maxSnapshotSize  int
	cfg              SnapshotDownloaderConfig
}

func (s *SnapshotDownloader) downloadPair(ctx context.Context, snapshotInterval int, pair common.Pair, ex common.Exchange, outChan chan common.Snapshot, logger zerolog.Logger, retryCount *int) bool {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			fmt.Printf("panic: %v\n", err)
			return
		}
	}()
	switch s.snapshotMode {
	case common.SnapshotModeTime:
		return s.downloadPairTime(ctx, time.Duration(snapshotInterval), pair, ex, outChan, logger, retryCount)
	case common.SnapshotModeTicks:
		return s.downloadPairTicks(ctx, snapshotInterval, pair, ex, outChan, logger, retryCount)
	}
	return false
}

func (s *SnapshotDownloader) downloadPairTime(ctx context.Context, snapshotInterval time.Duration, pair common.Pair, ex common.Exchange, outChan chan common.Snapshot, logger zerolog.Logger, retryCount *int) bool {
	var (
		wsCtx context.Context
		err   error
	)
	if wsCtx, err = ex.SubscribeEvents(pair); err != nil {
		logger.Err(err).Msg("subscribe event error")
		//errs = append(errs, fmt.Sprintf("%s [%s]: %v\n", ex.Name(), pair.String(), err))
		return false
	}

	ticker := time.NewTicker(snapshotInterval)

	logger.Info().Msg("started")
	ob, err := ex.GetSnapshot(pair)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s [%s]: %v\n", ex.Name(), pair.String(), err)
		return false
	}

	if outChan != nil {
		outChan <- ob
	}
	for {
		select {
		case <-s.ctx.Done():
			logger.Warn().Str("worker", "SnapshotDownloader").Msg("worker stop")
			ticker.Stop()
			return true
		case <-wsCtx.Done():
			logger.Warn().Str("worker", "SnapshotDownloader").Msg("websocket disconnected... restart")
			ticker.Stop()
			return false
		case <-ticker.C:
			ob, err := ex.GetSnapshot(pair)
			if err != nil {
				log.Error().Str("exchange", ex.Name()).Str("pair", pair.String()).Err(err).Send()
				return false
			}
			*retryCount = 0
			if outChan != nil {
				outChan <- ob
			}
		}
	}
}

func (s *SnapshotDownloader) downloadPairTicks(ctx context.Context, snapshotInterval int, pair common.Pair, ex common.Exchange, outChan chan common.Snapshot, logger zerolog.Logger, retryCount *int) bool {
	var (
		wsCtx context.Context
		err   error
	)
	if wsCtx, err = ex.SubscribeEvents(pair); err != nil {
		logger.Err(err).Msg("subscribe event error")
		//errs = append(errs, fmt.Sprintf("%s [%s]: %v\n", ex.Name(), pair.String(), err))
		return false
	}

	logger.Info().Msg("started")
	ob, err := ex.GetSnapshot(pair)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s [%s]: %v\n", ex.Name(), pair.String(), err)
		return false
	}

	if outChan != nil {
		outChan <- ob
	}
	ch := ex.SubscribeUpdates(pair, 10)
	defer ex.UnsubscribeUpdates(pair)
	lastSnapshot := ob
	for {
		select {
		case <-s.ctx.Done():
			logger.Warn().Str("worker", "SnapshotDownloader").Msg("worker stop")
			return true
		case <-wsCtx.Done():
			logger.Warn().Str("worker", "SnapshotDownloader").Msg("websocket disconnected... restart")
			return false
		case <-ch:
			ob, err := ex.GetSnapshot(pair)
			if lastSnapshot.Equal(ob) {
				continue
			}
			if err != nil {
				log.Error().Str("exchange", ex.Name()).Str("pair", pair.String()).Err(err).Send()
				return false
			}
			*retryCount = 0
			if outChan != nil {
				outChan <- ob
			}
		}
	}
}

func (s *SnapshotDownloader) Init(enabledExchanges ...string) error {
	s.exchanges = exchanges.GetExchanges(enabledExchanges...)
	if s.snapshotInterval == 0 {
		s.snapshotInterval = int(DefaultSnapshotInterval)
	}
	s.ctx, s.cancelFunc = context.WithCancel(context.Background())
	errs := []string{}
	for _, ex := range s.exchanges {
		if s.maxSnapshotSize != 0 {
			ex.SetMaxSnapshotSize(s.maxSnapshotSize)
		}
		if err := ex.Init(nil); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", ex.Name(), err))
		}
	}
	if len(errs) == 0 {
		fmt.Printf("Interval: %ds\n", time.Duration(s.snapshotInterval)/time.Second)
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
func (s *SnapshotDownloader) Run(ctx context.Context) error {
	errs := []string{}
	s.ctx, s.cancelFunc = context.WithCancel(context.Background())
	for _, ex := range s.exchanges {
		for _, pair := range s.pairs {
			logger := log.With().Str("exchange", ex.Name()).Str("pair", pair.String()).Logger()
			if !ex.PairAvailable(pair) {
				logger.Warn().Msg("pair not available")
				continue
			}
			go func(ctx context.Context, snapshotInterval int, pair common.Pair, ex common.Exchange, outChan chan common.Snapshot, logger zerolog.Logger) {
				retryCount := 0
				currentRetryInterval := retryInterval
			begin:
				if ok := s.downloadPair(ctx, snapshotInterval, pair, ex, outChan, logger, &retryCount); !ok && retryCount < maxRetryCount {
					retryCount++
					logger.Warn().Int("retry_count", retryCount).Msg("retry connection")
					time.Sleep(currentRetryInterval)
					currentRetryInterval += time.Second
					goto begin
				}
				s.cancelFunc()
				if s.cfg.ExitOnClose {
					os.Exit(1)
				}
			}(s.ctx, s.snapshotInterval, pair, ex, s.outChan, logger)
		}
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				_ = s.Close()
				return
			}
		}
	}()
	return utils.JoinErrsS(errs...)
}
func (s *SnapshotDownloader) Stop() {
	if s.cancelFunc != nil {
		s.cancelFunc()
	}
}
func (s *SnapshotDownloader) Close() error {
	if s.cancelFunc != nil {
		s.cancelFunc()
	}
	//close(s.outChan)
	errs := []string{}
	for _, ex := range s.exchanges {
		if err := ex.Close(); err != nil {
			errs = append(errs, fmt.Sprintf("%s: %v", ex.Name(), err))
		}
	}
	return utils.JoinErrsS(errs...)
}
