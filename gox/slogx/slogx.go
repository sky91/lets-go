package slogx

import (
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"strings"
)

func UseSlogFromEnv() error {
	lv := slog.LevelInfo
	if envSlogLevel := os.Getenv("SLOG_LEVEL"); envSlogLevel != "" {
		if err := lv.UnmarshalText([]byte(envSlogLevel)); err != nil {
			return errors.Wrapf(err, "unknown slog level: %s", envSlogLevel)
		}
	}

	handlerType := HandlerTypeText
	switch envSlogHandlerType := strings.ToLower(os.Getenv("SLOG_HANDLER_TYPE")); envSlogHandlerType {
	case "":
	case "text":
		handlerType = HandlerTypeText
	case "json":
		handlerType = HandlerTypeJSON
	default:
		return fmt.Errorf("unknown handler type: %s", envSlogHandlerType)
	}

	if err := UseSlog(lv, handlerType); err != nil {
		return errors.Wrap(err, "UseSlog() error")
	}
	return nil
}

func UseSlog(lv slog.Level, handlerType HandlerType) error {
	handlerOpts := slog.HandlerOptions{
		AddSource: true,
		Level:     lv,
	}
	var handler slog.Handler
	switch handlerType {
	case HandlerTypeText:
		handler = slog.NewTextHandler(os.Stderr, &handlerOpts)
	case HandlerTypeJSON:
		handler = slog.NewJSONHandler(os.Stderr, &handlerOpts)
	default:
		return fmt.Errorf("unknown handler type: %d", handlerType)
	}
	slog.SetDefault(slog.New(handler))
	slog.SetLogLoggerLevel(slog.LevelWarn)
	return nil
}

type HandlerType int

const (
	HandlerTypeText HandlerType = iota
	HandlerTypeJSON
)
