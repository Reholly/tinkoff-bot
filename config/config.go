package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Tinkoff TinkoffConfig `yaml:"tinkoff"`
}

type TinkoffConfig struct {
	SandboxApiToken string `yaml:"sandbox_api_token"`
}

func LoadConfig(configPath string) (Config, error) {
	config := Config{}
	file, err := os.ReadFile(configPath)

	if err != nil {
		return Config{}, errors.Wrap(err, "[ Config ] could not read file")
	}
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		return Config{}, errors.Wrap(err, "[ Config ] could not marshal config")
	}

	return Config{}, nil
}
