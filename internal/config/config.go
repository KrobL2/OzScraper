package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host         string `yaml:"host"`
	Storage_path string `yaml:"storage_path"`
	Batch_size   int    `yaml:"batch_size"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	data, err := os.ReadFile("./configs/main.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
