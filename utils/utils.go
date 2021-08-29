package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/niccoloCastelli/orderbooks/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var TimeProfileEnabled bool

func init() {
	TimeProfileEnabled = os.Getenv("TIME_PROFILE") != ""
}

func JoinErrs(errs ...error) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	errStrs := make([]string, len(errs))
	for i, e := range errs {
		errStrs[i] = e.Error()
	}
	return errors.New(strings.Join(errStrs, "\n"))
}
func JoinErrsS(errs ...string) error {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
func TimeProfile(name string, fn func()) {
	if !TimeProfileEnabled {
		fn()
		return
	}
	start := time.Now()
	fn()
	elapsed := time.Since(start).Nanoseconds()
	log.Info().Msgf("%s took %s", name, FormatDuration(time.Duration(elapsed)))
}
func FormatDuration(d time.Duration) string {
	if d > time.Second*10 {
		return fmt.Sprintf("%d s", d/time.Second)
	} else if d > time.Millisecond*10 {
		return fmt.Sprintf("%d ms", d/time.Millisecond)
	} else if d > time.Microsecond*10 {
		return fmt.Sprintf("%d μs", d/time.Microsecond)
	}
	return fmt.Sprintf("%d ns", d)
}

const EPSILON float64 = 0.00000001

func FloatEquals(a, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

func GetExchPairLogger(ex common.Exchange, pair common.Pair) zerolog.Logger {
	return log.Logger.With().Str("exchange", ex.Name()).Str("pair", pair.String()).Logger()
}

func UnmarshalResponse(resp *http.Response, out interface{}) error {
	var (
		rawData []byte
		err     error
	)
	defer resp.Body.Close()
	if rawData, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	if err := json.Unmarshal(rawData, out); err != nil {
		return err
	}
	return nil
}

func MakeExitChan() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	return c
}

func EnvOrDefault(env string, defaultVal string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return defaultVal
}
