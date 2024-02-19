package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	AppConf AppConfig
)

type AppConfig struct {
	DebugMode bool
	Port      string
	DBConfig  Database
}

func Load() {
	// set config based on env
	loadEnvVar()
	loadDBConfig("perqara_mysql")

	AppConf.DebugMode = viper.GetBool("debug_mode")
	AppConf.Port = viper.GetString("port")
}

func loadEnvVar() {
	var AppPath string

	// Bind OS environment variable
	viper.SetEnvPrefix("app")
	viper.BindEnv("env")
	viper.BindEnv("path")

	viper.SetConfigName("config")
	dir, _ := os.Getwd()
	AppPath = dir

	viper.SetConfigType("json")
	viper.AddConfigPath(AppPath + "/env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
