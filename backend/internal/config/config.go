package config

import (
	"fmt"
	"os"
)

type Config struct {
	FmpAPIKey string
	Port      string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("FMP_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("FMP_API_KEY env variable not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{FmpAPIKey: apiKey, Port: port}, nil
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf(":%s", c.Port)
}
