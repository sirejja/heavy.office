package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func Init(dev bool) {
	logger = New(dev)
}

func New(devel bool) *zap.Logger {
	var l *zap.Logger
	var err error
	if devel {
		l, err = zap.NewDevelopment()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = false
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		l, err = cfg.Build()
	}
	if err != nil {
		panic(err)
	}

	return l
}

func Get() *zap.Logger {
	return logger
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}
