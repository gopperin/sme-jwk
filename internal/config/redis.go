package config

import ()

// Redis Redis
type Redis struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
