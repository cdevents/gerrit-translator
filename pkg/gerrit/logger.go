package gerrit

import (
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// InitLogger Initialize a new production zap logger
func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic("failed to initialize logger Sync: " + err.Error())
		}
	}(logger)
	sugar = logger.Sugar()
}

// Log returns the zap logger If initialized
func Log() *zap.SugaredLogger {
	if sugar == nil {
		panic("zap logger is not initialized")
	}
	return sugar
}
