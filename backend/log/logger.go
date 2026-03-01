package log

import (
	"log/slog"
	"os"
)

func New(file *os.File) *slog.Logger {

	// handler options
	hOps := &slog.HandlerOptions{
		AddSource: true,
	}
	// "custom log handler"
	logger := slog.New(slog.NewJSONHandler(file, hOps))
	slog.SetDefault(logger)

	return logger
}
