package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Address string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
	} `mapstructure:"server"`
	Redis struct {
		Address  string `mapstructure:"address"`
		Password string `mapstructure:"password"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"redis"`
	OriginHost string
	IsClear    bool
	OriginPort int
}

func GetConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read the config file")
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config")
	}

	return config, nil
}
