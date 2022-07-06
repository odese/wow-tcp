package server

import (
	"testing"

	"wow/server/pkg/infrastructure/config"
	// log "wow/server/pkg/infrastructure/logging"
)

func TestRun(t *testing.T) {
	// log.LoggerSetup()
	config.Init()
	Run()
}
