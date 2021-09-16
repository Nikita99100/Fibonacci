package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Application struct {
	RestPort        string
	GrpcPort        string
	MemcacheAddress string
}

func Parse(filepath string) (*Application, error) {
	setDefaults()

	// Parse the file
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}
	// Unmarshal the config
	var cfg Application
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	// Validate the provided configuration
	return &cfg, nil
}
func setDefaults() {
	viper.SetDefault("RestPort", "80")
	viper.SetDefault("GrpcPort", "8080")
	viper.SetDefault("MemcacheAddress", "127.0.0.1:11211")
}
