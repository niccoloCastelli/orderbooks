package downloaders

import "context"

type Downloader interface {
	Init(enabledExchanges ...string) error
	Run(ctx context.Context) error
	Stop()
	Close() error
}
