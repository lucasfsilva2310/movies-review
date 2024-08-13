package Configuration

import "os"

type Config struct {
	PostgresURI string
}

func LoadPostgresConfig() *Config {
	return &Config{
		PostgresURI: os.Getenv("POSTGRES_URI"),
	}
}
