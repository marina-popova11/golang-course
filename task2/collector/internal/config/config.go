package config

import (
	"os"
	"time"
)

type Config struct {
	GRPC   GRPCConfig   `yaml:"grpc"`
	GitHub GitHubConfig `yaml:"github"`
	Logger LoggerConfig `yaml:"logger"`
}

type GRPCConfig struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type GitHubConfig struct {
	Token string `yaml:"token"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

func Load() (*Config, error) {
	return &Config{
		GRPC: GRPCConfig{
			Port:         ":8081",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		GitHub: GitHubConfig{
			Token: os.Getenv("GITHUB_TOKEN"),
		},
	}, nil
}
