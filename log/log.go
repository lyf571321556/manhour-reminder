package log

import (
	"github.com/lyf571321556/manhour-reminder/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

func InitLog() (err error) {
	logLevel := zap.InfoLevel
	if config.AppConfig.Debug {
		logLevel = zap.DebugLevel
	}
	zap.NewDevelopmentConfig()
	zapConfig := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: config.AppConfig.Debug,
		OutputPaths: []string{"stdout", config.AppConfig.LogPath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	Logger, err = zapConfig.Build()
	if err != nil {
		panic(err)
	}
	return err
}

func Debug(msg string, fields ...zap.Field) {
	if Logger == nil {
		return
	}
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	if Logger == nil {
		return
	}
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	if Logger == nil {
		return
	}
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	if Logger == nil {
		return
	}
	Logger.Fatal(msg, fields...)
}
