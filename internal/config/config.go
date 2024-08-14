package Configuration

import "os"

type Config struct {
	DatabaseUrl string
	Port        string
}

type DatabaseConfig struct {
	DatabaseUrl string
}

func LoadEnvConfig() *Config {
	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}
