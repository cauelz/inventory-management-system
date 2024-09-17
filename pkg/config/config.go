package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)


type Config struct {
	ServerAddress string
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
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
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
		DBUser: viper.GetString("POSTGRES_USER"),
		DBPassword: viper.GetString("POSTGRES_PASSWORD"),
		DBName: viper.GetString("POSTGRES_DB"),
	}

	return config, nil
}