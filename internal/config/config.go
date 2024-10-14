package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DSN string `yaml:"dsn"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
