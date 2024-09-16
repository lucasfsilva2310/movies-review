package Configuration

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DatabaseUrl string
	Port        string
	SecretKey   string
	AdminKey    string
}

type DatabaseConfig struct {
	DatabaseUrl string
}

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func LoadEnvConfig() *Config {
	return &Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),
		AdminKey:    os.Getenv("ADMIN_KEY"),
	}
}

type ApiConnection struct {
	Api *gin.Engine
}

func CreateApiConnection() *gin.Engine {
	apiConnection := gin.Default()
	apiConnection.SetTrustedProxies(nil)

	return apiConnection
}
