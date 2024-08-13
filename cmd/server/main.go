package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres Driver

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	db "github.com/lucasfsilva2310/movies-review/pkg/database"
)

func main() {
	// Load Environment Variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	config := Configuration.LoadEnvConfig()

	PORT := config.Port
	POSTGRES_URI := config.PostgresURI

	// Connect to DB
	dbConn, err := db.ConnectDB(&Configuration.DatabaseConfig{
		PostgresURI: POSTGRES_URI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Connect to API
	apiConnection := gin.Default()

	// Repositories

	// Services

	// Endpoints
	apiConnection.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	apiConnection.Run(":" + PORT)
}
