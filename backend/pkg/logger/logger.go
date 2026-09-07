package logger

import "go.uber.org/zap"

// L is the process-wide sugared logger instance.
var L *zap.SugaredLogger

// Init configures the global logger. Call once at startup.
// "development" gets human-readable console output + debug level;
// anything else gets JSON output tuned for log aggregators.
func Init(env string) {
	var base *zap.Logger
	var err error

	if env == "development" {
		base, err = zap.NewDevelopment()
	} else {
		base, err = zap.NewProduction()
	}
	if err != nil {
		panic("logger: failed to initialize zap: " + err.Error())
	}

	L = base.Sugar()
}

// Sync flushes any buffered log entries. Call via defer in main().
func Sync() {
	if L != nil {
		_ = L.Sync()
	}
}
