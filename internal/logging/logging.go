package logging

import (
	"io"
	"log/slog"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

type slogLogger struct {
	logger *slog.Logger
}

func getLevel(s string) slog.Leveler {
	switch s {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	}
	return slog.LevelInfo
}

func NewJsonLogger(w io.Writer, level string) Logger {
	return &slogLogger{
		logger: slog.New(
			slog.NewJSONHandler(w, &slog.HandlerOptions{
				Level: getLevel(level),
			}),
		),
	}
}

func NewTextLogger(w io.Writer, level string) Logger {
	return &slogLogger{
		logger: slog.New(
			slog.NewTextHandler(w, &slog.HandlerOptions{
				Level: getLevel(level),
			}),
		),
	}
}

func (l *slogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *slogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *slogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *slogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}
