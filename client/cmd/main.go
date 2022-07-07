package main

import (
	"wow/client/pkg/client"
	"wow/client/pkg/infrastructure/config"
	log "wow/client/pkg/infrastructure/logger"
)

func init() {
	log.LoggerSetup()
	config.Init()
}

func main() {
	client.Run()
}
