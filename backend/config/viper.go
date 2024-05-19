package config

import (
	"github.com/spf13/viper"
)

// config stores all configuration of the application
// the value are read by viper from a config file or enviroment variable
type Config struct {
	CacheTimeExpiration int    `mapstructure:"CACHE_TIME_EXPIRATION"`
	CacheExpiration     bool   `mapstructure:"CACHE_EXPIRATION"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	Capacity            int    `mapstructure:"CAPACITY"`
	RedisAddress        string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword       string `mapstructure:"REDIS_PASSWORD"`
	RedisDbName         int    `mapstructure:"REDIS_DB_NAME"`
	SwaggerUrl          string `mapstructure:"SWAGGER_URL"`
	SwaggerConfigPath   string `mapstructure:"SWAGGER_CONFIG_PATH"`
}

// LoadConfig reads configuration from file or environment variable.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
