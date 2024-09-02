package zerolog

import (
	"os"
	"time"

	"github.com/leonzag/treport/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ logger.Logger = new(zeroLogger)

type zeroLogger struct {
	base zerolog.Logger
}

func NewLogger() (*zeroLogger, error) {
	if os.Getenv("APP_ENV") == "development" {
		return NewLoggerDevelop(), nil
	}
	return NewLoggerProduction(), nil
}

func NewLoggerProduction() *zeroLogger {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	log.Logger = log.Logger.Level(zerolog.InfoLevel)

	return &zeroLogger{base: log.Logger}
}

func NewLoggerDevelop() *zeroLogger {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	})
	log.Logger = log.Logger.Level(zerolog.TraceLevel)
	log.Logger = log.Logger.With().Caller().Logger()

	return &zeroLogger{base: log.Logger}
}

func (l *zeroLogger) Debugf(template string, args ...any) {
	l.base.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Infof(template string, args ...any) {
	l.base.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warnf(template string, args ...any) {
	l.base.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Errorf(template string, args ...any) {
	l.base.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatalf(template string, args ...any) {
	l.base.Fatal().Msgf(template, args...)
}
