package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const cmdRoot = "core"

// Setup 载入配置文件
func Setup(path string) {

	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	Application = &ApplicationStruct{
		Port: viper.GetString("application.port"),
	}

	CORS.Enable = viper.GetBool("cors.enable")
	CORS.AllowOrigins = viper.GetStringSlice("cors.AllowOrigins")
	CORS.AllowMethods = viper.GetStringSlice("cors.AllowMethods")
	CORS.AllowHeaders = viper.GetStringSlice("cors.AllowHeaders")
	CORS.AllowCredentials = viper.GetBool("cors.AllowCredentials")
	CORS.MaxAge = viper.GetInt("cors.MaxAge")

}
