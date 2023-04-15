package config

import (
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string
	Port        int
	DebugMode   bool
}

func NewConfig() *Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	debugMode, _ := strconv.ParseBool(os.Getenv("DEBUG_MODE"))

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        port,
		DebugMode:   debugMode,
	}
}
