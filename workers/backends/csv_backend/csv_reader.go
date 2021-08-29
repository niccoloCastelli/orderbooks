package csv_backend

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/common/constants"
	"github.com/niccoloCastelli/orderbooks/data_formats/ob_format"
	"github.com/niccoloCastelli/orderbooks/orderbook"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/niccoloCastelli/orderbooks/workers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"io"
	"os"
	"path"
	"regexp"
	"sync"
	"time"
)

type snapshotFileReader struct {
	path string
	csv  *csv.Reader
}

func NewSnapshotReader(baseFs afero.Fs, logger *zerolog.Logger) (workers.SnapshotReader, error) {
	if logger == nil {
		logger = &log.Logger
	}
	readerLogger := logger.With().Str("reader", "csv").Logger()
	return (&snapshotCsvReader{baseFs: baseFs, logger: readerLogger}).Init()
}

type snapshotCsvReader struct {
	baseFs   afero.Fs
	inChan   chan common.Snapshot
	ctx      context.Context
	cancelFn context.CancelFunc
	//openFiles     map[string]*snapshotFileReader
	existingFiles map[string]map[string][]time.Time // existingFiles[exchangeName][pair][file_t0, file_t1, file_t2,...]
	logger        zerolog.Logger
	wg            sync.WaitGroup
	cachedRows    map[string][][]string
}

func (s *snapshotCsvReader) Init() (workers.SnapshotReader, error) {
	if err := s.scanExchanges(); err != nil {
		return nil, err
	}
	return s, nil
}
func (s *snapshotCsvReader) Close() error {
	if s.cancelFn != nil {
		s.cancelFn()
		s.cancelFn = nil
	}
	s.wg.Wait()
	/*for _, f := range s.openFiles {
		if err := f.reader.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "file close err: %v", err)
		}
	}*/
	return nil
}
func (s *snapshotCsvReader) Exchanges() []string {
	exchangesList := make([]string, 0, len(s.existingFiles))
	for exc, _ := range s.existingFiles {
		exchangesList = append(exchangesList, exc)
	}
	return exchangesList
}
func (s *snapshotCsvReader) Pairs(exchangeName string) []common.Pair {
	if exchange, ok := s.existingFiles[exchangeName]; ok {
		pairs := make([]common.Pair, 0, len(exchange))
		for p, _ := range exchange {
			pair, err := common.PairsFromString(p)
			if err != nil {
				s.logger.Err(err).Str("exchange", exchangeName).Str("pair", p).Send()
				continue
			}
			pairs = append(pairs, *pair)
		}
		return pairs
	}
	return []common.Pair{}
}
func (s *snapshotCsvReader) TimeRange(exchangeName string, pair common.Pair) *utils.TimeRange {
	if exchange, ok := s.existingFiles[exchangeName]; ok {
		if timeRange, ok := exchange[pair.String()]; ok {
			return &utils.TimeRange{
				Start: timeRange[0],
				End:   timeRange[len(timeRange)-1],
			}
		}
	}
	return nil
}
func (s *snapshotCsvReader) Read(exchangeName string, pair common.Pair, timeRange utils.TimeRange, interval int64, mode common.SnapshotMode, size int, ctx context.Context) (chan common.Snapshot, error) {
	ch := make(chan common.Snapshot, 1)
	var (
		intervalD = time.Duration(interval)
	)
	if size == 0 {
		size = 50
	}
	if exchange, ok := s.existingFiles[exchangeName]; ok {
		if dayFiles, ok := exchange[pair.String()]; ok {
			go func(ch chan common.Snapshot) {
				var (
					logger = s.logger.With().Str("interval", intervalD.String()).Str("exchange", exchangeName).Str("pair", pair.String()).Int("num_orders", size).Logger()
					ob     *orderbook.OrderBook
					err    error
				)

				defer close(ch)
				for _, dayFile := range dayFiles {
					select {
					case <-ctx.Done():
						return
					default:
						nextDay := dayFile.AddDate(0, 0, 1)
						dayLogger := logger.With().Time("day", dayFile).Logger()
						//fmt.Println(dayFile, nextDay, dayFile.Before(timeRange.Start), timeRange.Start.Before(nextDay))
						if (dayFile.Before(timeRange.Start) || dayFile.Equal(timeRange.Start)) && timeRange.Start.Before(nextDay) {
							ob, err = s.readDailySnapshot(ob, timeRange, dayFile, exchangeName, pair, interval, mode, size, dayLogger, ch)
							if err != nil {
								return
							}
						} else if dayFile.After(timeRange.Start) && dayFile.Before(timeRange.End) {
							ob, err = s.readDailySnapshot(ob, timeRange, dayFile, exchangeName, pair, interval, mode, size, dayLogger, ch)
							if err != nil {
								return
							}
						} else if dayFile.After(timeRange.End) {
							dayLogger.Debug().Msg("channel closed")
							return
						}
					}

				}
			}(ch)
			return ch, nil
		}
		return nil, errors.New("pair not found")
	}
	return nil, errors.New("exchange not found")
}

func (s *snapshotCsvReader) readDailySnapshot(ob *orderbook.OrderBook, timeRange utils.TimeRange, dayFile time.Time, exchangeName string, pair common.Pair, interval int64, mode common.SnapshotMode, size int, logger zerolog.Logger, ch chan common.Snapshot) (*orderbook.OrderBook, error) {
	var (
		lastSnapshotTime time.Time
		lastSnapshot     common.Snapshot
		initTime         time.Time
		intervalD        = time.Duration(interval)
		tickCount        = 0
	)
	if ob == nil {
		ob = orderbook.NewOrderBook(logger, nil).LimitSnapshotSize(size)
		if err := s.readFirstSnapshot(ob, dayFile, exchangeName, pair, logger); err != nil {
			return nil, err
		}
	}

	lastSnapshotTime = ob.LastUpdate()
	f, err := s.readFile(constants.FileTypeEvents, dayFile, exchangeName, pair)
	if err != nil {
		logger.Err(err).Stack().Msg("read file error")
		return nil, err
	}
	/*if rows, err = f.csv.ReadAll(); err != nil {
		logger.Err(err).Msg("parse header error")
		return nil, err
	}*/
	if _, err := f.csv.Read(); err != nil {
		return nil, err
	}
	for {
		row, err := f.csv.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Err(err).Interface("row", row).Send()
			return nil, err
		}
		if len(row) < 4 {
			logger.Warn().Strs("row", row).Msg("skip bad row")
			continue
		}
		parsedRow, err := ob_format.ParseObRow(row)
		if err != nil {
			logger.Err(err).Msg("parse row error")
			return nil, err
		}
		evt := parsedRow.ToEvent()
		if evt.Type == common.EventTypeInit {
			logger.Warn().Str("time", evt.Timestamp.String()).Interface("event", evt).Msg("init event")
			ob = orderbook.NewOrderBook(logger, nil).LimitSnapshotSize(size)
			initTime = evt.Timestamp
			continue
		}
		if evt.Type == common.EventTypeAdd && (evt.Order.Timestamp.Equal(initTime) || evt.Order.Timestamp.Sub(initTime) < time.Second*10) { // &&
			logger.Trace().Str("time", evt.Timestamp.String()).Time("timestamp", evt.Timestamp).Interface("event", evt).Msg("skip init event")
			ob.Update(evt.Order, true)
			continue
		}
		ob.Update(evt.Order, false)
		tickCount++
		if mode == common.SnapshotModeTime {
			if ob.LastUpdate().Sub(lastSnapshotTime) >= intervalD && (ob.LastUpdate().Equal(timeRange.Start) || ob.LastUpdate().After(timeRange.Start)) {
				snapshot := ob.GetSnapshot()
				lastSnapshotTime = snapshot.Timestamp
				ch <- snapshot
			}
		} else {
			if tickCount >= int(interval) {
				snapshot := ob.GetSnapshot()
				if !lastSnapshot.Equal(snapshot) {
					lastSnapshotTime = snapshot.Timestamp
					ch <- snapshot
					tickCount = 0
					lastSnapshot = snapshot
				}

			}
		}

		if ob.LastUpdate().Equal(timeRange.End) || ob.LastUpdate().After(timeRange.End) {
			logger.Debug().Msg("complete")
			return ob, nil
		}
	}
	return ob, nil
}
func (s *snapshotCsvReader) readFirstSnapshot(ob *orderbook.OrderBook, day time.Time, exchangeName string, pair common.Pair, logger zerolog.Logger) error {
	var (
		err                 error
		initialSnapshotTime time.Time
	)
	if ob == nil {
		return errors.New("no order book")
	}
	f, err := s.readFile(constants.FileTypeSnapshots, day, exchangeName, pair)
	if err != nil {
		logger.Err(err).Stack().Msg("read file error")
		return err
	}
	if _, err := f.csv.Read(); err != nil {
		return err
	}
	for {
		row, err := f.csv.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Err(err).Interface("row", row).Send()
			return err
		}
		parsedRow, err := ob_format.ParseObRow(row)
		if err != nil {
			logger.Err(err).Interface("row", row).Msg("parse row error")
			return err
		}
		order := parsedRow.ToOrder()
		if initialSnapshotTime.IsZero() {
			initialSnapshotTime = order.Timestamp
		}
		if order.Timestamp.After(initialSnapshotTime) {
			break
		}
		ob.Update(order, true)
	}
	logger.Debug().Msg("order book initialized")
	return nil
}
func (s *snapshotCsvReader) readFile(fileType constants.FileType, ts time.Time, exchange string, pair common.Pair) (*snapshotFileReader, error) {
	var (
		filePath = getCurrentFilePath(fileType, ts, exchange)
		fileName = pair.String() + ".csv"
		fullPath = path.Join(filePath, fileName)
		//fileKey   = getFileKey(fileType, exchange, pair)
		csvReader io.Reader
	)
	if !fileExists(s.baseFs, fullPath) {
		if !fileExists(s.baseFs, fullPath+".gz") {
			s.logger.Debug().Str("path", fullPath).Msg("file not found")
			return nil, os.ErrNotExist
		}
		buf := bytes.NewBuffer(nil)
		gz, err := s.baseFs.OpenFile(fullPath+".gz", os.O_RDONLY, 0640)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		stat, err := gz.Stat()
		if err != nil {
			return nil, err
		}
		if err := utils.GunzipWrite(buf, gz); err != nil {
			return nil, err
		}
		csvReader = buf
		s.logger.Debug().Str("path", fullPath).Int64("size", stat.Size()).Msg("gzip extracted")
	} else {
		f, err := s.baseFs.OpenFile(fullPath, os.O_RDONLY, 0640)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		csvReader = f
	}
	fileReader := &snapshotFileReader{
		path: fullPath,
		csv:  csv.NewReader(csvReader),
	}
	fileReader.csv.FieldsPerRecord = -1
	return fileReader, nil
}
func (s *snapshotCsvReader) scanExchanges() error {
	s.existingFiles = map[string]map[string][]time.Time{}
	reg, err := regexp.Compile(EventsFileMatchRe)
	if err != nil {
		return err
	}
	err = afero.Walk(s.baseFs, EventsBasePath, func(fp string, info os.FileInfo, err error) error {
		if !info.IsDir() && reg.MatchString(fp) {
			m := reg.FindStringSubmatch(fp)
			exchange := m[2]
			fileDay, err := time.Parse("20060102", m[3]+m[4]+m[5])
			if err != nil {
				return err
			}
			pair, err := common.PairsFromString(m[6])
			if err != nil {
				return err
			}
			if _, ok := s.existingFiles[exchange]; !ok {
				s.existingFiles[exchange] = map[string][]time.Time{}
			}
			if _, ok := s.existingFiles[exchange][pair.String()]; !ok {
				s.existingFiles[exchange][pair.String()] = []time.Time{}
			}
			days := s.existingFiles[exchange][pair.String()]
			added := false
			for idx, day := range days {
				if day.After(fileDay) {
					numDays := len(days)
					before := days[0:idx]
					after := days[idx:numDays]
					days = append(before, fileDay)
					days = append(days, after...)
					added = true
				}
			}
			if !added {
				days = append(days, fileDay)
			}
			s.existingFiles[exchange][pair.String()] = days
		}
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Println(s.Exchanges())
	return nil
}
