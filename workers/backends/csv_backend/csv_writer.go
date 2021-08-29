package csv_backend

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/common/constants"
	"github.com/niccoloCastelli/orderbooks/workers"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type snapshotFileWriter struct {
	path   string
	writer io.WriteCloser
	csv    *csv.Writer
}

func NewSnapshotWriter(baseFs afero.Fs, inChan chan common.Snapshot, saveEvents bool) workers.SnapshotWriter {
	return (&snapshotCsvWriter{baseFs: baseFs, openFiles: map[string]*snapshotFileWriter{}, saveEvents: saveEvents}).Init(inChan)
}

type snapshotCsvWriter struct {
	baseFs     afero.Fs
	inChan     chan common.Snapshot
	ctx        context.Context
	cancelFn   context.CancelFunc
	openFiles  map[string]*snapshotFileWriter
	saveEvents bool
	wg         sync.WaitGroup
}

func (s *snapshotCsvWriter) fileExists(filePath string) bool {
	if exists, _ := afero.Exists(s.baseFs, filePath); !exists {
		return false
	}
	if isDir, _ := afero.IsDir(s.baseFs, filePath); isDir {
		return false
	}
	return true
}
func (s *snapshotCsvWriter) getCurrentFile(fileType constants.FileType, ts time.Time, exchange string, pair common.Pair) (*snapshotFileWriter, error) {
	filePath := getCurrentFilePath(fileType, ts, exchange)
	fileName := pair.String() + ".csv"
	fullPath := path.Join(filePath, fileName)
	fileKey := getFileKey(fileType, exchange, pair)
	if openFile, ok := s.openFiles[fileKey]; ok {
		if openFile.path == fullPath {
			return openFile, nil
		}
		if err := openFile.writer.Close(); err != nil {
			return nil, err
		}
		delete(s.openFiles, fileKey)
	}
	newFile := snapshotFileWriter{
		path:   fullPath,
		writer: nil,
	}
	if exists, _ := afero.DirExists(s.baseFs, filePath); !exists {
		if err := s.baseFs.MkdirAll(filePath, 0750); err != nil {
			return nil, err
		}
	}
	fileExisting := fileExists(s.baseFs, fullPath)
	f, err := s.baseFs.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return nil, err
	}
	newFile.csv = csv.NewWriter(f)
	newFile.writer = f
	if !fileExisting {
		switch fileType {
		case constants.FileTypeSnapshots:
			if err := newFile.csv.Write((&common.Order{}).Headers()); err != nil {
				return nil, err
			}
		case constants.FileTypeEvents:
			if err := newFile.csv.Write((&common.Event{}).Headers()); err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("unknown file type")
		}

	}
	s.openFiles[fileKey] = &newFile

	return &newFile, nil
}
func (s *snapshotCsvWriter) Init(ch chan common.Snapshot) workers.SnapshotWriter {
	s.inChan = ch
	return s
}
func (s *snapshotCsvWriter) Run(ctx context.Context) error {
	s.ctx, s.cancelFn = context.WithCancel(ctx)
	if !s.saveEvents {
		log.Info().Msg("ignore events")
	}
	go func(inChan chan common.Snapshot, ctx context.Context, wg *sync.WaitGroup) {
		wg.Add(1)
		for {
			select {
			case sn := <-s.inChan:
				logger := log.With().Str("exchange", sn.Exchange).Str("pair", sn.Pair.String()).Logger()
				if sn.Timestamp.IsZero() {
					logger.Error().Str("snapshot", fmt.Sprintf("%.30v", sn)).Msg("timestamp zero")
					break
				}
				ordersFile, err := s.getCurrentFile(constants.FileTypeSnapshots, sn.Timestamp, sn.Exchange, sn.Pair)
				if err != nil {
					logger.Err(err).Msg("open file error")
					break
				}
				for _, o := range sn.Orders {
					if err := ordersFile.csv.Write(o.Row()); err != nil {
						logger.Err(err).Str("row", strings.Join(o.Row(), ",")).Msg("write row error")
					}
				}
				ordersFile.csv.Flush()
				if err := ordersFile.csv.Error(); err != nil {
					logger.Err(err).Msg("write file error")
				}

				if s.saveEvents {
					eventsFile, err := s.getCurrentFile(constants.FileTypeEvents, sn.Timestamp, sn.Exchange, sn.Pair)
					if err != nil {
						logger.Err(err).Msg("open file error")
						break
					}
					for _, o := range sn.Events {
						if err := eventsFile.csv.Write(o.Row()); err != nil {
							logger.Err(err).Str("row", strings.Join(o.Row(), ",")).Msg("write row error")
						}
					}
					eventsFile.csv.Flush()
					if err := eventsFile.csv.Error(); err != nil {
						logger.Err(err).Msg("write file error")
					}
				}
				continue
			case <-s.ctx.Done():
				wg.Done()
				return
			}
		}
	}(s.inChan, s.ctx, &s.wg)
	return nil
}
func (s *snapshotCsvWriter) Close() error {
	if s.cancelFn != nil {
		s.cancelFn()
		s.cancelFn = nil
	}
	s.wg.Wait()
	for _, f := range s.openFiles {
		if err := f.writer.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "file close err: %v", err)
		}
	}
	return nil
}
