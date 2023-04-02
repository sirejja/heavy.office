package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Kafka struct {
		Brokers []string `yaml:"brokers"`
		Topics  struct {
			OrderStatus struct {
				Topic string `yaml:"topic"`
			} `yaml:"order-status"`
		} `yaml:"topics"`
		GroupName       string `yaml:"group-name"`
		BalanceStrategy string `yaml:"balance-strategy"`
	} `yaml:"kafka"`
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
