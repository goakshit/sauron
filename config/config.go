package config

import (
	"os"
	"strconv"

	"github.com/goakshit/sauron/internal/types"
)

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	env := getEnv(key, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}
	return result
}

func getPostgresConf() types.PostgresConf {
	return types.PostgresConf{
		DatabaseName: getEnv("POSTGRES_DB_NAME", "sdb"),
		User:         getEnv("POSTGRES_USER", "dharakshit"),
		Passwd:       getEnv("POSTGRES_PASSWD", "passwd"),
		Host:         getEnv("POSTGRES_HOST", "localhost"),
		Port:         getEnvAsInt("POSTGRES_PORT", 5432),
		SSLMode:      getEnv("POSTGRES_SSLMODE", "disable"),
		Timezone:     getEnv("POSTGRES_TZ", "UTC"),
	}
}

// New - Initialize Configuration
func New() types.Config {
	return types.Config{
		AppPort:  getEnvAsInt("APP_PORT", 80),
		AppEnv:   getEnv("APP_ENV", "local"),
		Postgres: getPostgresConf(),
	}
}
