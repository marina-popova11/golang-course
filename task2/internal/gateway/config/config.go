package config

import (
	"os"
	"time"
)

type Config struct {
	HTTP      HTTPConfig      `yaml:"http"`
	Collector CollectorConfig `yaml:"collector"`
	Swagger   SwaggerConfig   `yaml:"swagger"`
}

type HTTPConfig struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type CollectorConfig struct {
	Address string `yaml:"address"`
}

type SwaggerConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func Load() (*Config, error) {
	return &Config{
		HTTP: HTTPConfig{
			Port:         ":8081",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Collector: CollectorConfig{
			Address: os.Getenv("GITHUB_TOKEN"),
		},
		Swagger: SwaggerConfig{
			Enabled: true,
			Path:    "/swagger",
		},
	}, nil
}
