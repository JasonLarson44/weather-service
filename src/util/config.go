package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	OWBaseUri string `mapstructure:"OW_BASE_URI"`
	Port      string `mapstructure:"PORT"`
	ApiKey    string `mapstructure:"OW_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config %v", err)
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Failed to unmarshal config %v", err)
	}
	return
}
