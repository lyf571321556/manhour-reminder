package log

import (
	"github.com/lyf571321556/manhour-reminder/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

func InitLog() (err error) {
	logLevel := zap.InfoLevel
	outputPaths := make([]string, 0)
	if conf.AppConfig.Debug {
		logLevel = zap.DebugLevel
		outputPaths = append(outputPaths, "stdout")
	} else {
		outputPaths = append(outputPaths, conf.AppConfig.LogPath)
	}
	zap.NewDevelopmentConfig()
	zapConfig := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: conf.AppConfig.Debug,
		OutputPaths: outputPaths,
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
