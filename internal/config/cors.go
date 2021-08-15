package config

import ()

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool     `mapstructure:"enable" json:"enable" yaml:"enable"`
	AllowOrigins     []string `mapstructure:"AllowOrigins" json:"AllowOrigins" yaml:"AllowOrigins"`
	AllowMethods     []string `mapstructure:"AllowMethods" json:"AllowMethods" yaml:"AllowMethods"`
	AllowHeaders     []string `mapstructure:"AllowHeaders" json:"AllowHeaders" yaml:"AllowHeaders"`
	AllowCredentials bool     `mapstructure:"AllowCredentials" json:"AllowCredentials" yaml:"AllowCredentials"`
	MaxAge           int      `mapstructure:"MaxAge" json:"MaxAge" yaml:"MaxAge"`
}
