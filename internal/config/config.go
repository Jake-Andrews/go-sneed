package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
    Port   string `env:"PORT" envDefault:"5151"`
    PG_URI string `env:"PG_URI,required"`
}

func LoadConfig() *Config {
    var cfg Config
    if err := env.Parse(&cfg); err != nil {
        log.Fatalf("Error loading config: %v", err)
    }
    log.Print(cfg)
    return &cfg
}
