package config

import ()

// CORSStruct 跨域请求配置参数
type CORSStruct struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

// CORS CORS
var CORS = new(CORSStruct)
