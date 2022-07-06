package config

import (
	"testing"

	// "wow/server/configs"
)

func TestInit(t *testing.T) {
	Init()
}

func TestGetConfig(t *testing.T) {
	Init()
	GetConfig()
}