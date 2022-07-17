package zap

import (
	"brewing_support/app/utils/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger = create()

// GetLogger ロガーの取得関数
func GetLogger() *zap.Logger {
	return logger
}

func create() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/application.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		zap.DebugLevel,
	)

	if config.Mode() == "production" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		)
	}

	if config.Mode() == "test" {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			w,
			zap.DebugLevel,
		)
	}

	return zap.New(core)
}
