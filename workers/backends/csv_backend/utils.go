package csv_backend

import (
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/niccoloCastelli/orderbooks/common/constants"
	"github.com/spf13/afero"
	"time"
)

const (
	EventsBasePath    = "events"
	EventsFileMatchRe = `^(\w+)\/(.*?)\/(\d{4})\/(\d{2})\/(\d{2})\/([A-Z]+-[A-Z]+)\.csv\.gz$`
)

func getCurrentFilePath(fileType constants.FileType, ts time.Time, exchange string) string {
	return fmt.Sprintf("/%s/%s/%04d/%02d/%02d", fileType, exchange, ts.Year(), ts.Month(), ts.Day())
}
func getFileKey(fileType constants.FileType, exchange string, pair common.Pair) string {
	return fmt.Sprintf("%s:%s:%s", fileType, exchange, pair.String())
}
func fileExists(fs afero.Fs, filePath string) bool {
	if exists, _ := afero.Exists(fs, filePath); !exists {
		return false
	}
	if isDir, _ := afero.IsDir(fs, filePath); isDir {
		return false
	}
	return true
}
