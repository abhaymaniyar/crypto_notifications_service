package test_helpers

import (
	"service_dashboard/internal/logger"
	"sync"
)

var loggerInitOnce sync.Once

func InitializeLogger() {
	loggerInitOnce.Do(func() {
		logger.Init(logger.DEBUG)
	})
}
