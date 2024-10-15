package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DSN string `yaml:"dsn"`
}

func MustLoad(configPath string) *Config {
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	return &config
}
