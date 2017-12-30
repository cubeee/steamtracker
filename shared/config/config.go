package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(module string, env *string) error {
	viper.SetConfigName(module + "-config." + *env)
	viper.AddConfigPath("./resources/config/")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
