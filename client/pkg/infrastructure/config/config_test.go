package config

import (
	"testing"

	"wow/server/pkg/utils"
)

func TestInit(t *testing.T) {
	utils.InitiateTestConfig()
}

func TestGetConfig(t *testing.T) {
	exptectedNetwork := "tcp"
	exptectedAddress := "localhost:6090"

	got := utils.InitiateTestConfig()

	if got.GetString("server.network") != exptectedNetwork {
		t.Errorf("Expected %v, got %v", exptectedNetwork, got.GetString("server.network"))
	}

	if got.GetString("server.address") != exptectedAddress {
		t.Errorf("Expected %v, got %v", exptectedAddress, got.GetString("server.address"))
	}
}

func TestGetConfigName(t *testing.T) {
	expected := "config"
	got := GetConfigName()
	if expected != got {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
