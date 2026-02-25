package log

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {

	// handler options
	hOps := &slog.HandlerOptions{
		AddSource: true,
	}
	// "custom log handler"
	logger := slog.New(slog.NewJSONHandler(os.Stderr, hOps))
	slog.SetDefault(logger)

	return logger
}
