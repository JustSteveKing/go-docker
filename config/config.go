package config

import "os"

// AppConfig is the Application configuration struct
type AppConfig struct {
	Name    string
	Version string
}

// DatabaseConfig is the Database configuration struct
type DatabaseConfig struct {
	User       string
	Password   string
	Host       string
	Name       string
	Connection string
}

// Config is the Configuration struct
type Config struct {
	Database DatabaseConfig
	App      AppConfig
}

// New returns a new Config Struct
func New() *Config {
	return &Config{
		App: AppConfig{
			Name:    env("APP_NAME", "Go App"),
			Version: env("APP_VERSION", "v1.0"),
		},
		Database: DatabaseConfig{
			User:       env("DB_USER", ""),
			Password:   env("DB_PASSWORD", ""),
			Host:       env("DB_HOST", ""),
			Name:       env("DB_NAME", ""),
			Connection: env("DB_CONNECTION", "tcp"),
		},
	}
}

// env is a simple helper function to read an environment variable or return a default value
func env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
