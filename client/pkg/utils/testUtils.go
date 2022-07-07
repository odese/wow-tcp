package utils

import (
	log "wow/server/pkg/infrastructure/logger"

	"github.com/spf13/viper"
)

func InitiateTestConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigType("yml")
	config.SetConfigName("config")
	config.AddConfigPath("C:/workspace/go/src/wow-tcp/server/configs")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("Error on reading config")
	}

	return config
}

// IsStringElementInArray checks if a string element is in a string array
func IsStringElementInArray(e string, arr []string) bool {
	for _, x := range arr {
		if x == e {
			return true
		}
	}
	return false
}
