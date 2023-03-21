package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Services struct {
		Loms struct {
			URL string `yaml:"url"`
		} `yaml:"loms"`
		Products struct {
			Token string `yaml:"token"`
			URL   string `yaml:"url"`
		} `yaml:"products-service"`
	} `yaml:"services"`
	Storage struct {
		PostgresDSN string
	}
}

func New() *ConfigStruct {
	return &ConfigStruct{}
}

func (c *ConfigStruct) Init() error {
	op := "ConfigStruct.Init"
	rawYAML, err := os.ReadFile("config.yaml")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = yaml.Unmarshal(rawYAML, &c)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_ = godotenv.Load()
	c.Storage.PostgresDSN = os.Getenv("DB_DSN")
	if c.Storage.PostgresDSN == "" {
		return fmt.Errorf("%s: %w", op, errors.New("Database credentials are not provided"))
	}

	return nil
}
