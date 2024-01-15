package config

import (
	"chat/logger"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Url      string
	DbName   string
	Username string
	Password string
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("Error reading config file, %s", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
