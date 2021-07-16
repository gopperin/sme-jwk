package config

import ()

// ApplicationStruct ApplicationStruct
type ApplicationStruct struct {
	Port string
}

// Application Application
var Application = new(ApplicationStruct)
