package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Theme struct {
		Color string `yaml:"color"`
	}
	Navigation struct {
		ShowingSplash bool `yaml:"showingSplash"`
	}
}

func loadConfig(path string) (*Config, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &config, nil
}

func AppConfig() *Config {
	config, err := loadConfig(".")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}
	return config
}
