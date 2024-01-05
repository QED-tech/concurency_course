package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	defaultEngineType     = "in_memory"
	defaultAddress        = "127.0.0.1:8080"
	defaultMaxConnections = 30
	defaultLogLevel       = "production"
	defaultLogOutput      = "stdout"
)

type EngineConfig struct {
	Type string `yaml:"type"`
}

type NetworkConfig struct {
	MaxConnections int    `yaml:"max_connections"`
	Address        string `yaml:"address"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

type Config struct {
	Engine  EngineConfig  `yaml:"engine"`
	Network NetworkConfig `yaml:"network"`
	Logging LoggingConfig `yaml:"logging"`
}

func ReadConfig() (*Config, error) {
	conf, err := os.ReadFile("./internal/database/config/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	c := defaultConfig()

	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &c, nil
}

func defaultConfig() Config {
	return Config{
		Engine: EngineConfig{
			Type: defaultEngineType,
		},
		Network: NetworkConfig{
			MaxConnections: defaultMaxConnections,
			Address:        defaultAddress,
		},
		Logging: LoggingConfig{
			Level:  defaultLogLevel,
			Output: defaultLogOutput,
		},
	}
}
