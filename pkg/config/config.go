package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)


type Config struct {
	ServerAddress string
}

func LoadConfig() (*Config, error) {
	
	err := gotenv.Load()

	if err != nil {
		log.Println("Could not load .env file", err)
	}

	viper.AutomaticEnv()

	viper.SetDefault("SERVER_ADDRESS", ":8080")

	config := &Config{
		ServerAddress: viper.GetString("SERVER_ADDRESS"),
	}

	return config, nil
}