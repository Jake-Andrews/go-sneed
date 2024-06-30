package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
    Server_port string `env:"SERVER_PORT" envDefault:"5151"`
    Server_host string `env:"SERVER_HOST" envDefault:"localhost"`
    PG_URI      string `env:"PG_URI,required"`
}

func LoadConfig() *Config {
    var cfg Config
    if err := env.Parse(&cfg); err != nil {
        log.Fatalf("Error loading config: %v", err)
    }
    return &cfg
}
