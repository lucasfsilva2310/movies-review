package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres Driver

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	db "github.com/lucasfsilva2310/movies-review/pkg/database"
)

func main() {
	// Load Environment Variables
	if os.Getenv("LOAD_ENV_FILE") == "true" {

		err := godotenv.Load("../../.env")
		if err != nil {
			log.Fatal(err)
		}
	}
	config := Configuration.LoadEnvConfig()

	PORT := config.Port
	DATABASE_URL := config.DatabaseUrl

	// Connect to DB
	log.Println(DATABASE_URL)
	dbConn, err := db.ConnectDB(&Configuration.DatabaseConfig{
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	// Connect to API
	apiConnection := gin.Default()
	apiConnection.SetTrustedProxies(nil)
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
