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

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetUInt64(key string) uint64 {
	return uint64(GetInt64(key))
}
