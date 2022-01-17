package core

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

//FSLocation defines sources and destinations for filesync
type FSLocation struct {
	Host string `mapstructure:"host"`
	Path string `mapstructure:"path"`
	Type string `mapstructure:"type"`
}

//SyncEntry for fssync
type SyncEntry struct {
	Source FSLocation `mapstructure:"source"`
	Target FSLocation `mapstructure:"target"`
}

// Config object for fssync
type Config struct {
	Version string      `mapstructure:"version"`
	Entries []SyncEntry `mapstructure:"entries"`
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
