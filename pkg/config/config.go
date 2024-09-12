package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Directory    string `mapstructure:"directory"`
	TimeInterval uint   `mapstructure:"time_interval"`
}

// LoadConfig loads config from file or override it with the command line arguments
func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
