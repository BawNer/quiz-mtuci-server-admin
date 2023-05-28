package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
}

// Logger -.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn().Int("severity", 400)
}

func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error().Int("severity", 500)
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal().Int("severity", 800)
}
