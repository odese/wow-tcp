package utils

import (
	log "wow/server/pkg/infrastructure/logger"

	"github.com/spf13/viper"
)

func InitiateTestConfig() {
	config := viper.New()
	config.SetConfigType("yml")
	config.SetConfigName("config")
	config.AddConfigPath("C:/workspace/go/src/wow-tcp/server/configs")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("Error on reading config")
	}
}
