package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/lpernett/godotenv"
)

type Config struct {
	AppName string `envconfig:"APP_NAME" default:"calculator-api"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
