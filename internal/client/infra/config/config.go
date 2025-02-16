package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseFile string
	GrpcAddress  string
}

func NewConfig() Config {
	if os.Getenv("IS_TEST_ENV") == "true" {
		_ = godotenv.Load("../../../.env.client")
	} else {
		_ = godotenv.Load(".env.client")
	}

	return Config{
		DatabaseFile: getEnv("DATABASE_FILE", ""),
		GrpcAddress:  getEnv("GRPC_ADDRESS", ""),
	}
}

func getEnvInt(key string, def int) int {
	v, e := strconv.Atoi(getEnv(key, strconv.Itoa(def)))
	if e != nil {
		return def
	} else {
		return v
	}
}

func getEnv(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return def
}
