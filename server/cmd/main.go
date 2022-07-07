package main

import (
	"wow/server/pkg/infrastructure/config"
	log "wow/server/pkg/infrastructure/logger"
	"wow/server/pkg/server"
)

func init() {
	log.LoggerSetup()
	config.Init()
}

func main() {
	server.Run()
}
