package config

import (
	"os"

	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

// Config represents the configuration
type Config struct {
	Verbose bool `default:"false"`
	Port    uint `default:"9090"`

	Prometheus struct {
		URL string
	}

	SQL struct {
		Connection string
	}
}

// LoadConfig Loads configuration from config file
func LoadConfig(configFile string) (*Config, error) {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		log.Warningf("Config file not found, %v. Running with default configuration", configFile)
	}

	cfg := Config{}
	err = configor.Load(&cfg, configFile)
	if err != nil {
		log.Errorf("Unable to parse config file, %v", configFile)
		return nil, err
	}

	return &cfg, nil
}