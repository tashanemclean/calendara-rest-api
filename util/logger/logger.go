package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

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

func addAttrs(pipeSeparatedTags string, err error) []slog.Attr {
	var attrsGroup []slog.Attr
	tags := strings.Split(pipeSeparatedTags, "|")
	attrs := append(attrsGroup, slog.Any("tags", tags), slog.Any("error", err))
	return attrs
}

func Error(msg string, pipeSeparatedTags string, err error, args ...any) {
	if !logger.Enabled(context.Background(), slog.LevelError) {
		return
	}
	attrs := addAttrs(pipeSeparatedTags, err)
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pcs[0])
	r.Add(args...)
	_ = logger.Handler().WithAttrs(attrs).Handle(context.Background(), r)
}

func Debug(msg string, pipeSeparatedTags string, args ...any) {
	if !logger.Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var attrsGroup []slog.Attr
	tags := strings.Split(pipeSeparatedTags, "|")
	attrs := append(attrsGroup, slog.Any("tags", tags))
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelDebug, msg, pcs[0])
	r.Add(args...)
	_ = logger.Handler().WithAttrs(attrs).Handle(context.Background(), r)
}

func Info(msg string, pipeSeparatedTags string, err error, args ...any) {
	if !logger.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	attrs := addAttrs(pipeSeparatedTags, err)
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelInfo, msg, pcs[0])
	r.Add(args...)
	_ = logger.Handler().WithAttrs(attrs).Handle(context.Background(), r)
}

func Warn(msg string, pipeSeparatedTags string, err error, args ...any) {
	if !logger.Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	attrs := addAttrs(pipeSeparatedTags, err)
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelWarn, msg, pcs[0])
	r.Add(args...)
	_ = logger.Handler().WithAttrs(attrs).Handle(context.Background(), r)
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