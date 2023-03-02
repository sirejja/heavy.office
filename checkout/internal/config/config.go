package config

import (
	"fmt"
	"os"

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

	return nil
}
