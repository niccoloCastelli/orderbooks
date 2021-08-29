package writers

import (
	"context"
	"github.com/niccoloCastelli/orderbooks/common"
	orderbooks "github.com/niccoloCastelli/orderbooks/server/proto"
	"github.com/niccoloCastelli/orderbooks/workers"
	"github.com/olivere/elastic"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

type ElasticSearchWriterConfig struct {
	indexPrefix      string
	esChanBufferSize int
	esBulkSize       int
}

func NewElasticSearchWriterConfig(indexPrefix string, esChanBufferSize int, esBulkSize int) ElasticSearchWriterConfig {
	return ElasticSearchWriterConfig{indexPrefix: indexPrefix, esChanBufferSize: esChanBufferSize, esBulkSize: esBulkSize}
}

type ElasticSearchWriter struct {
	serverUrls []string
	inputChan  chan common.Snapshot
	client     *elastic.Client
	logger     zerolog.Logger
	exchange   string
	pair       common.Pair
	ctx        context.Context
	cancelFn   context.CancelFunc
	conf       ElasticSearchWriterConfig
}

func NewElasticSearchWriter(logger zerolog.Logger, exchange string, pair common.Pair, conf ElasticSearchWriterConfig, serverUrls ...string) *ElasticSearchWriter {
	return &ElasticSearchWriter{serverUrls: serverUrls, exchange: exchange, pair: pair, conf: conf, logger: logger.With().Str("component", "writer").Logger()}
}
func (e *ElasticSearchWriter) IndexName() string {
	return UniqueName(e.conf.indexPrefix, e.pair, e.exchange)
}
func (e *ElasticSearchWriter) Init(ch chan common.Snapshot) workers.SnapshotWriter {
	var err error
	e.client, err = elastic.NewClient(elastic.SetURL(e.serverUrls[0]))
	e.ctx, e.cancelFn = context.WithCancel(context.Background())
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	e.inputChan = ch
	return e
}

func (e *ElasticSearchWriter) Run(ctx context.Context) error {
	esChan := make(chan indexRequest, e.conf.esChanBufferSize)
	indexName := e.IndexName()
	go func() {
		for {
			select {
			case snapshot, ok := <-e.inputChan:
				if !ok {
					return
				}
				pbOrders := make([]*orderbooks.Event, len(snapshot.Orders))
				snapshotMap := map[string]interface{}{}
				for i, order := range snapshot.Orders {
					pbOrders[i] = &orderbooks.Event{
						Timestamp: order.Timestamp.UnixNano(),
						OrderSide: orderbooks.OrderSide(order.Side),
						Amount:    float32(order.Amount),
						Price:     float32(order.Price),
					}
					priceIdx := strconv.Itoa(i * 2)
					volumeIdx := strconv.Itoa((i * 2) + 1)
					snapshotMap[priceIdx] = order.Price
					snapshotMap[volumeIdx] = order.Amount
				}
				snapshotMap["ts"] = snapshot.Timestamp.Format(time.RFC3339)

				esChan <- indexRequest{
					esRequest: elastic.NewBulkIndexRequest().Id(snapshot.Timestamp.String()).Doc(snapshotMap),
					ts:        snapshot.Timestamp,
				}
				//.Id("").BodyJson().Do(ctx)
			case <-e.ctx.Done():
				close(esChan)
				return
			case <-ctx.Done():
				close(esChan)
				return
			}
		}
	}()
	go func() {
		requests := make([]elastic.BulkableRequest, 0, e.conf.esBulkSize)
		counter := 0
		for {
			select {
			case req, ok := <-esChan:
				var snapshotTime time.Time
				if ok {
					requests = append(requests, req.esRequest)
					snapshotTime = req.ts
				}
				if len(requests) == e.conf.esBulkSize || (!ok && len(requests) > 0) {
					if resp, err := e.client.Bulk().Index(indexName).Add(requests...).Do(ctx); err != nil {
						e.logger.Fatal().Err(err).Send()
						return
					} else if resp.Errors {
						e.logger.Fatal().
							Bool("errors", resp.Errors).
							Int("took", resp.Took).
							Interface("items", resp.Items).
							Int("total", counter).
							Time("snapshot_time", snapshotTime).
							Msg("errors")
					} else {
						counter += len(requests)
						e.logger.Debug().
							Bool("errors", resp.Errors).
							Int("took", resp.Took).
							Int("items", len(resp.Items)).
							Int("total", counter).
							Time("snapshot_time", snapshotTime).
							Msg("request OK")
					}
					requests = make([]elastic.BulkableRequest, 0, 1000)
				}
			case <-ctx.Done():

				return
			case <-e.ctx.Done():

				return
			}
		}
	}()
	return nil
}

func (e *ElasticSearchWriter) Close() error {
	e.cancelFn()
	return nil
}

type indexRequest struct {
	esRequest *elastic.BulkIndexRequest
	ts        time.Time
}
