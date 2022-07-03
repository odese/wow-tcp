package main

import (
	"wow/server/pkg/base/config"
	log "wow/server/pkg/base/logging"
	"wow/server/pkg/server"
)

func init() {
	log.LoggerSetup()
	config.Init()
}

func main() {
	server.Run()
}
