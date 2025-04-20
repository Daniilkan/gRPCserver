package config

import (
	"OrderProject/pkg/postgres"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.Config

	GRPCPort int `yaml:"GRPC_PORT" env:"GRPC_PORT" env-default:"50051"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
