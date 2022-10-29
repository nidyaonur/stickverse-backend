package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	ApplicationHost string
	ApplicationPort string
	Debug           bool
	BasePrefix      string
	DatabaseUrl     string
	WriterDBUrl     string
	ReaderDBUrl     string
	JwtSecret       string
	AutoMigrate     bool
}

// ParseEnv Get environment value from os
// If an environment required and not set raises a panic
func ParseEnv(key string, required bool, dft string) string {
	_ = godotenv.Load()
	value := os.Getenv(key)
	if value == "" && required {
		panic(fmt.Sprintf("Environment variable not found: %v", key))
	} else if value == "" {
		return dft
	}
	return value
}

//FIXME: Remove unnecessary variables before deployment
func GetEnvironment() *Environment {
	env := &Environment{
		ApplicationHost: ParseEnv("APPLICATION_HOST", false, "0.0.0.0"),
		ApplicationPort: ParseEnv("APPLICATION_PORT", false, "8000"),
		Debug:           ParseEnv("DEBUG", false, "false") == "true",
		DatabaseUrl:     ParseEnv("DB_URL", false, "postgres://root:secret@localhost:5432/stickverse"),
		WriterDBUrl:     ParseEnv("DB_URL", false, ""),
		ReaderDBUrl:     ParseEnv("DB_URL", false, ""),
		JwtSecret:       ParseEnv("JWT_SECRET", false, "test_secret"),
		AutoMigrate:     ParseEnv("AUTO_MIGRATE", false, "false") == "true",
	}
	if env.WriterDBUrl == "" {
		env.WriterDBUrl = env.DatabaseUrl
	}
	if env.ReaderDBUrl == "" {
		env.ReaderDBUrl = env.DatabaseUrl
	}
	return env
}
