package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct{}

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
