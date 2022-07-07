package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {

}

func TestGetConfig(t *testing.T) {
	exptectedNetwork := "tcp"
	exptectedAddress := "localhost:6090"

	got := viper.New()
	got.SetConfigType("yml")
	got.SetConfigName(GetConfigName())
	got.AddConfigPath("C:/workspace/go/src/wow-tcp/server/configs")
	err := got.ReadInConfig()
	if err != nil {
		t.Error(err.Error())
	}

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
