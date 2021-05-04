package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

const (
	AppName = "APP_NAME"

	DbName     = "DB_%v_DBNAME"
	DbUser     = "DB_%v_USER"
	DbPassword = "DB_%v_PWD"
	DbHost     = "DB_%v_HOST"
	DbPort     = "DB_%v_PORT"

	RedisActive = "REDIS_ACTIVE"
	RedisHost   = "REDIS_HOST"
	RedisPort   = "REDIS_PORT"

	ElasticHost   = "ELASTIC_HOST"
	ElasticUser   = "ELASTIC_USER"
	ElasticPwd    = "ELASTIC_PWD"
	ElasticScheme = "ELASTIC_SCHEME"

	LogType     = "LOG_TYPE"
	LogFilePath = "LOG_FILE_PATH"
)

func SetConfigFile(name, path, extension string) {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	viper.SetConfigType(extension)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
}

func GetString(key string) string {
	return viper.GetString(fmt.Sprintf("%v", key))
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(fmt.Sprint(key))
}

func GetStringIfExists(key string, defaultValue string) string {
	val := viper.GetString(fmt.Sprintf("%v", key))
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

func GetInt(key string) int {
	return viper.GetInt(fmt.Sprint(key))
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(fmt.Sprint(key))
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetDBDriver() string {
	return viper.GetString("database.driver")
}

func GetConfigUsed(osEnv string, localEnv string) string {
	configVariable := os.Getenv(osEnv)
	if configVariable == "" {
		configVariable = viper.GetString(localEnv)
	}
	return configVariable
}

func GetAppName() string {
	return GetConfigUsed(AppName, "name")
}

func GetDBName(driver string) string {
	return GetConfigUsed(fmt.Sprintf(DbName, driver), fmt.Sprintf("database.%v.dbname", driver))
}

func GetDBUser(driver string) string {
	return GetConfigUsed(fmt.Sprintf(DbUser, driver), fmt.Sprintf("database.%v.username", driver))
}

func GetDBPass(driver string) string {
	return GetConfigUsed(fmt.Sprintf(DbPassword, driver), fmt.Sprintf("database.%v.password", driver))
}

func GetDBHost(driver string) string {
	return GetConfigUsed(fmt.Sprintf(DbHost, driver), fmt.Sprintf("database.%v.host", driver))
}

func GetDBPort(driver string) string {
	return GetConfigUsed(fmt.Sprintf(DbPort, driver), fmt.Sprintf("database.%v.port", driver))
}

func IsUsingRedis() bool {
	config := GetConfigUsed(RedisActive, "cache.redis.active")
	b, err := strconv.ParseBool(config)
	if err != nil {
		log.Fatal("Error reading redis active", err)
	}
	return b
}

func GetRedisAddr() string {
	addrRedis := GetConfigUsed(RedisHost, "cache.redis.address")
	portRedis := GetConfigUsed(RedisPort, "cache.redis.port")
	return fmt.Sprintf("%v:%v", addrRedis, portRedis)
}

func GetElasticHost() string {
	return GetConfigUsed(ElasticHost, "elasticsearch.host")
}

func GetElasticUsername() string {
	return GetConfigUsed(ElasticUser, "elasticsearch.username")
}

func GetElasticPassword() string {
	return GetConfigUsed(ElasticPwd, "elasticsearch.password")
}

func GetElasticScheme() string {
	return GetConfigUsed(ElasticScheme, "elasticsearch.scheme")
}

func GetLogType() string {
	return GetConfigUsed(LogType, "logType")
}

func GetLogFile() string {
	return GetConfigUsed(LogFilePath, "logFile.path")
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(fmt.Sprint(key))
}

func GetStringMapString(key string) []byte {
	v := viper.GetStringMapString(key)
	vByte, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Can't marshaling configs")
	}

	return vByte
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(fmt.Sprintf("%v", key))
}
