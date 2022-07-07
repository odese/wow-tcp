package main

import (
	"testing"

	"wow/server/pkg/infrastructure/config"
	log "wow/server/pkg/infrastructure/logger"
	"wow/server/pkg/server"
)

func TestMain(t *testing.T) {

	log.LoggerSetup()
	config.Init()
	server.Run()
}
