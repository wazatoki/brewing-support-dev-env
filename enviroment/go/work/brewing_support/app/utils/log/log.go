package log

import (
	"brewing_support/app/infrastructures/zap"
)

// Debug output debug message
func Debug(message string) {
	zap.GetLogger().Debug(message)
}

// Info output infomation message
func Info(message string) {
	zap.GetLogger().Info(message)
}

// Warn output warning message
func Warn(message string) {
	zap.GetLogger().Warn(message)
}

// Error output error message
func Error(message string) {
	zap.GetLogger().Error(message)
}
