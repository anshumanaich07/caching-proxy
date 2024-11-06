package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		Db      string `mapstructure:"db"`
		Retries int    `mapstructure:"retries"`
	} `mapstructure:"database"`
}

func GetConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read the config file")
	}

	// unmarshal
	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config")
	}

	return config, nil
}
