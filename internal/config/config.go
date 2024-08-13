package Configuration

import "os"

type Config struct {
	PostgresURI string
	Port        string
}

type DatabaseConfig struct {
	PostgresURI string
}

func LoadEnvConfig() *Config {
	return &Config{
		PostgresURI: os.Getenv("POSTGRES_URI"),
		Port:        os.Getenv("PORT"),
	}
}
