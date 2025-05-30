package configs

import (
	"github.com/caarlos0/env/v10"
)

// Config holds the configuration for the application.
type Config struct {
	Port     int  `env:"K8SESE_PORT" envDefault:"8080"`
	Debug    bool `env:"K8SESE_DEBUG" envDefault:"false"`
	Interval int  `env:"K8SESE_INTERVAL" envDefault:"10"` // in seconds
}

// LoadConfig loads the configuration from environment variables using the caarlos0/env package.
func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
