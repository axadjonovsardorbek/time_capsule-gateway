package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	GATEWAY_PORT string
	MEMORY_PORT string
	TIMELINE_PORT string

	MEMORY_HOST string
	TIMELINE_HOST string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.GATEWAY_PORT = cast.ToString(coalesce("GATEWAY_PORT", ":8081"))
	config.MEMORY_PORT = cast.ToString(coalesce("MEMORY_PORT", ":50051"))
	config.TIMELINE_PORT = cast.ToString(coalesce("TIMELINE_PORT", ":50052"))
	config.MEMORY_HOST = cast.ToString(coalesce("MEMORY_HOST", "memory-service"))
	config.TIMELINE_HOST = cast.ToString(coalesce("TIMELINE_HOST", "timeline-service"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
