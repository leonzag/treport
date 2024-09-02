package zap

import (
	"os"
	"time"

	"github.com/leonzag/treport/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ logger.Logger = new(zapLogger)

type zapLogger struct {
	base *zap.SugaredLogger
}

func NewLogger() (*zapLogger, error) {
	if os.Getenv("APP_ENV") == "development" {
		return NewLoggerDevelop()
	}
	return NewLoggerProduction()
}

func NewLoggerProduction() (*zapLogger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	zapConfig.EncoderConfig.TimeKey = "time"
	l, err := zapConfig.Build()
	logger := l.Sugar()
	return &zapLogger{base: logger}, err
}

func NewLoggerDevelop() (*zapLogger, error) {
	zapConfig := zap.NewDevelopmentConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	zapConfig.EncoderConfig.TimeKey = "time"
	l, err := zapConfig.Build()
	logger := l.Sugar()
	return &zapLogger{base: logger}, err
}

func (l *zapLogger) Debugf(template string, args ...any) {
	l.base.Debugf(template, args...)
}

func (l *zapLogger) Infof(template string, args ...any) {
	l.base.Infof(template, args...)
}

func (l *zapLogger) Warnf(template string, args ...any) {
	l.base.Warnf(template, args...)
}

func (l *zapLogger) Errorf(template string, args ...any) {
	l.base.Errorf(template, args...)
}

func (l *zapLogger) Fatalf(template string, args ...any) {
	l.base.Fatalf(template, args...)
}
