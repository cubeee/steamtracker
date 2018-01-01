package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(module string, env *string) error {
	viper.SetConfigName("config." + *env)
	viper.AddConfigPath("./resources/config/" + module + "/")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
