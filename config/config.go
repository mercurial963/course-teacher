package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

const (
	databaseUrlKey = "DATABASE_URL"
	listenAddrKey  = "LISTEN_ADDR"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func init() {
	var err error
	config = viper.New()
	config.AutomaticEnv()
	if _, err = os.Stat("./config/development.yaml"); !errors.Is(err, os.ErrNotExist) {
		config.SetConfigType("yaml")
		config.SetConfigName("development")
		config.AddConfigPath("../config")
		config.AddConfigPath("config/")
		err = config.ReadInConfig()
		if err != nil {
			log.Fatal("error on parsing configuration file")
		}
	}

	// config.SetDefault(databaseUrlKey, "postgres://postgres:mypass@127.0.0.1:5432/class?sslmode=disable")
}

func GetDatabaseURL() string { return config.GetString(databaseUrlKey) }
func GetListenAddr() string  { return config.GetString(listenAddrKey) }
