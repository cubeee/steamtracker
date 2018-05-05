package config

import (
	"github.com/spf13/viper"
)

func ReadConfig(module string, env string) error {
	viper.SetConfigFile(GetConfigFilePath(module, env))
	return viper.ReadInConfig()
}

func GetConfigFilePath(module string, env string) string {
	return "./resources/config/" + module + "/config." + env + ".yml"
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
