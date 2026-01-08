package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APPEnv   string
	APPPort  int
	DBHost   string
	DBPort   int
	DBUser   string
	DBPass   string
	DBName   string
	LOGLevel string
}

func loadenv() {
	if err := godotenv.Load(".env"); err != nil {
		godotenv.Load(".env.development")
	}
}

func loadEnvWithKey(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func loadEnvIntWithKey(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		var intVal int
		_, err := fmt.Sscanf(val, "%d", &intVal)
		if err == nil {
			return intVal
		}
	}
	return fallback
}

func NewConfig() *Config {
	loadenv()
	return &Config{
		APPEnv:   loadEnvWithKey("APP_ENV", "development"),
		APPPort:  loadEnvIntWithKey("APP_PORT", 8080),
		DBHost:   loadEnvWithKey("DB_HOST", "localhost"),
		DBPort:   loadEnvIntWithKey("DB_PORT", 5433),
		DBUser:   loadEnvWithKey("DB_USER", "kamamuchi"),
		DBPass:   loadEnvWithKey("DB_PASS", "supersecretpassword_dev"),
		DBName:   loadEnvWithKey("DB_NAME", "onlydrive_dev"),
		LOGLevel: loadEnvWithKey("LOG_LEVEL", "debug"),
	}
}
