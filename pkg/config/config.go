package config

import (
	"fmt"
	"github.com/orangbus/m3d/pkg/debug"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

func LoadConfig() {
	viper = viperlib.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		debug.Panic("配置读取错误", err)
	}
	viper.WatchConfig()
}

func GetMysqlUrl() string {
	host := getEnvString("db.host", "localhost")
	port := getEnvInt("db.port", 3306)
	database := getEnvString("db.database", "cloud_movie")
	user_name := getEnvString("db.user", "root")
	password := getEnvString("db.password", "")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user_name, password, host, port, database)
	return dsn
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func getEnvString(key string, defVal ...string) string {
	v := viper.GetString(key)
	if v == "" && defVal != nil {
		v = defVal[0]
	}
	return v
}
func getEnvInt(key string, defVal ...int) int {
	v := viper.GetInt(key)
	if v == 0 && defVal != nil {
		v = defVal[0]
	}
	return v
}
