package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func Initialize() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func getEnvIsInt(key string, defaultValue int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}
