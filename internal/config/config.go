package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	Server struct {
    Host   string 
    Port   int    
    Debug  bool  
    Secret string
	}
}

// LoadConfig loads the configuration from a configuration file or environment variables.
func LoadConfig() (*Config, error) {
	// Load the configuration from a file, if present
	viper.SetConfigName("config")
  viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("Failed to read configuration file: %v", err)
		}
	}

	// Unmarshal the configuration into a struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal configuration: %v", err)
	}

	return &cfg, nil
}

