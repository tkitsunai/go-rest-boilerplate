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
	Logger zerolog.Logger
)

func init() {
	once.Do(func() {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%-6s]", i))
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf(" %s ", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatCaller = func(i interface{}) string {
			t := fmt.Sprintf("%s", i)
			s := strings.Split(t, ":")
			if 2 != len(s) {
				return t
			}
			f := filepath.Base(s[0])
			return f + ":" + s[1]
		}
		Logger = zerolog.New(output).Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
	})
}
