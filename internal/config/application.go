package config

import ()

// Application Application
type Application struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}
