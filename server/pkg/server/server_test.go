package server

import (
	"testing"

	"wow/server/pkg/infrastructure/config"
	// "wow/server/pkg/infrastructure/logger"
)

func TestRun(t *testing.T) {
	// log.LoggerSetup()
	config.Init()
	Run()
}
