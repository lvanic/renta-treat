package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL   string
	Port          string
	DebugMode     bool
	Secret        string
	TokenLifespan int
}

var once sync.Once
var instance *Config

func NewConfig() *Config {
	once.Do(func() {
		port := viper.Get("PORT").(string)
		databaseURL := viper.Get("DatabaseURL").(string)
		debugMode := viper.GetBool("DEBUG_MODE")
		secret := viper.Get("API_SECRET").(string)
		tokenLifespan := viper.GetInt("TOKEN_HOUR_LIFESPAN")
		instance = &Config{
			DatabaseURL:   databaseURL,
			Port:          port,
			DebugMode:     debugMode,
			Secret:        secret,
			TokenLifespan: tokenLifespan,
		}
	})
	return instance

}
