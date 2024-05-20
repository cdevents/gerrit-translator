package gerrit

import (
	"errors"
	"syscall"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// init Initialize a new production zap logger
func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer func(logger *zap.Logger) {
		err := Sync(sugar)
		if err != nil {
			panic("failed to initialize logger Sync: " + err.Error())
		}
	}(logger)
	sugar = logger.Sugar()
}

func Sync(logger *zap.SugaredLogger) error {
	err := logger.Sync()
	if !errors.Is(err, syscall.EINVAL) {
		return err
	}
	return nil
}

// Log returns the zap logger If initialized
func Log() *zap.SugaredLogger {
	if sugar == nil {
		panic("zap logger is not initialized")
	}
	return sugar
}
