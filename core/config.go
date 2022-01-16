package core

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config object for fssync
type Config struct {
	Version string `mapstructure:"version"`
}

// NewConfig generates Config object pointer for config file
func NewConfig(configFile string) (*Config, error) {
	// Get abs path of config file
	absPath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, err
	}

	// Set viper props for config file
	viper.SetConfigName(filepath.Base(absPath))
	viper.AddConfigPath(filepath.Dir(absPath))
	viper.SetConfigType(strings.TrimPrefix(filepath.Ext(absPath), "."))

	// Read in viper configs
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Parse viper configs into struct
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
