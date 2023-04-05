package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Web struct {
		Port string `yaml:"port"`
	} `yaml:"web"`
	Kafka struct {
		Brokers []string `yaml:"brokers"`
		Topics  struct {
			OrderStatus string `yaml:"order-status"`
		} `yaml:"topics"`
	} `yaml:"kafka"`
	Storage struct {
		PostgresDSN string
	}
	CancelOrdersCronPeriod string
	OutboxCronPeriod       string
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

	c.CancelOrdersCronPeriod = os.Getenv("CANCEL_ORDERS_CRON_SCHEDULE")
	if c.CancelOrdersCronPeriod == "" {
		return fmt.Errorf("%s: %w", op, errors.New("Cancel orders cron not scheduled"))
	}

	c.OutboxCronPeriod = os.Getenv("OUTBOX_CRON_SCHEDULE")
	if c.OutboxCronPeriod == "" {
		return fmt.Errorf("%s: %w", op, errors.New("Outbox cron not scheduled"))
	}

	return nil
}
