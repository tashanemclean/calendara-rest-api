package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)
var logger *slog.Logger

func SetupLogger() {
	var level = slog.LevelDebug
	level = slog.Level(GetLogLevel())

	var options = &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Rename source key to match Logging Definition
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}

			return a
		},
	}

	handler := slog.NewJSONHandler(os.Stdout, options)
	logger = slog.New(handler.WithAttrs([]slog.Attr{slog.Any("caller", "calendara-rest-api-api")}))
}

func GetLogLevel() slog.Level {
	logLevel := strings.ToLower(viper.GetString("LOG_LEVEL"))
	switch logLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		fmt.Println("No Log Level Found Setting Log Level to Info.")
		return slog.LevelInfo
	}
}