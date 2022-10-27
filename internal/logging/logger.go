package logging

import (
	"github.com/jalavosus/slackthing/internal/utils"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger(name string, core zapcore.Core) *zap.Logger {
	loggerOpts := []zap.Option{
		zap.Fields(zap.String("name", name)),
	}

	if core != nil {
		return zap.New(
			core,
			loggerOpts...,
		)
	}

	l, _ := zap.NewProduction(loggerOpts...)

	return l
}

func NewLogger(name string) *zap.Logger {
	return newLogger(name, nil)
}

func NewDevLogger(name string) *zap.Logger {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.InfoLevel,
	)

	l := newLogger(name, core)

	return l
}

func NewLoggerFromEnv(name string) *zap.Logger {
	isDev := utils.IsDev()
	if isDev {
		return NewDevLogger(name)
	}

	return NewLogger(name)
}