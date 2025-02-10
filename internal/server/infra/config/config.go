package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBAddress   string
	GrpcAddress string
}

func NewConfig() Config {
	if os.Getenv("IS_TEST_ENV") == "true" {
		_ = godotenv.Load("../../../.env.server")
	} else {
		_ = godotenv.Load(".env.server")
	}

	return Config{
		DBAddress:   getEnv("DATABASE_DSN", ""),
		GrpcAddress: getEnv("GRPC_ADDRESS", ""),
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
