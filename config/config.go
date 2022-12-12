package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("env.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("讀取 env.yml 失敗: %s", err.Error())
	}
}

func GetEnvConfig() *viper.Viper {
	return viper.GetViper()
}
