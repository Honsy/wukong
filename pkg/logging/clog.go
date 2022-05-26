package logging

import "test/log"

var (
	logger log.Logger
)

func init() {
	logger = log.NewDefaultLogrusLogger().WithPrefix("WuKong")
	logger.SetLevel(log.DebugLevel)
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debug(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Info(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Warn(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.Error(v)
}

// Fatal output logs at fatal level
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.Fatal(v)
}

// Fatal output logs at fatal level
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}
