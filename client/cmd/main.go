package main

import (
	"wow/client/pkg/client"
	"wow/client/pkg/base/config"
	log "wow/client/pkg/base/logging"
)

func init() {
	log.LoggerSetup()
	config.Init()
}

func main() {
	client.Run()
}
