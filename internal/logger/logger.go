package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	once   = sync.Once{}
	logger zerolog.Logger
)

func init() {
	once.Do(func() {
		out := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
		out.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf(" %s ", i)
		}
		out.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		out.FormatCaller = func(i interface{}) string {
			t := fmt.Sprintf("%s", i)
			s := strings.Split(t, ":")
			if 2 != len(s) {
				return t
			}
			f := filepath.Base(s[0])
			return f + ":" + s[1]
		}
		logger = zerolog.New(out).Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	})
}

func Get() *zerolog.Logger {
	return &logger
}
