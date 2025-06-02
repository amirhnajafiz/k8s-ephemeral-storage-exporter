package configs

import (
	"fmt"

	"github.com/caarlos0/env/v10"
)

// Config holds the configuration for the application.
type Config struct {
	Port        int    `env:"K8SESE_PORT" envDefault:"8080"`
	Debug       bool   `env:"K8SESE_DEBUG" envDefault:"false"`
	Interval    int    `env:"K8SESE_INTERVAL" envDefault:"10"` // in seconds
	K8SLocalAPI string `env:"K8SESE_K8S_LOCAL_API" envDefault:""`
}

// LoadConfig loads the configuration from environment variables using the caarlos0/env package.
func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: %w", err)
	}

	return &cfg, nil
}
