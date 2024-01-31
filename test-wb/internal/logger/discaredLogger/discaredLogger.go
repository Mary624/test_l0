package discaredlogger

import (
	"context"
	"log/slog"
)

type discaredHandler struct{}

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

func NewDiscardHandler() *discaredHandler {
	return &discaredHandler{}
}

func (h *discaredHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *discaredHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *discaredHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *discaredHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
