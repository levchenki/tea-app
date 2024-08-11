package config

import (
	"context"
	"log"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Env string `env:"ENV, default=DEV"`
	HTTPServer
	Storage
}

type HTTPServer struct {
	Port        string        `env:"SERVER_PORT, required"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT, required"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT, required"`
}

type Storage struct {
	Host     string `env:"DB_HOST, required"`
	Port     int    `env:"DB_PORT, required"`
	Name     string `env:"DB_NAME, required"`
	Username string `env:"DB_USERNAME, required"`
	Password string `env:"DB_PASSWORD, required"`
}

const (
	EnvDevelopment = "DEV"
	EnvProduction  = "PROD"
)

func MustLoad() *Config {
	ctx := context.Background()
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Env != EnvDevelopment && cfg.Env != EnvProduction {
		log.Fatal("The environment can only be DEV or PROD.")
	}

	return &cfg
}
