package config

import (
	"os"
	"time"
)

type Config struct {
	HTTP      HTTPConfig      `yaml:"http" env:"http"`
	Collector CollectorConfig `yaml:"collector" env:"collector"`
	Swagger   SwaggerConfig   `yaml:"swagger"`
}

type HTTPConfig struct {
	Port         string        `yaml:"port" env:"PORT"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type CollectorConfig struct {
	Address string `yaml:"address" env:"COLLECTOR_ADDRESS"`
}

type SwaggerConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func Load() (*Config, error) {
	cfg := &Config{
		HTTP: HTTPConfig{
			Port:         ":8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Collector: CollectorConfig{
			Address: "localhost:8081",
		},
		Swagger: SwaggerConfig{
			Enabled: true,
			Path:    "/swagger",
		},
	}

	if envAddress := os.Getenv("COLLECTOR_ADDRESS"); envAddress != "" {
		cfg.Collector.Address = envAddress
	}

	if envPort := os.Getenv("PORT"); envPort != "" {
		cfg.HTTP.Port = envPort
	}

	return cfg, nil
}
