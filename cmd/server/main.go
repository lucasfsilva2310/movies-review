package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres Driver

	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/movies"
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
	movieRepo := movies.NewRepository(dbConn)

	// Services
	movieService := movies.NewService(movieRepo)

	// Endpoints

	// Health
	apiConnection.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Movies
	apiConnection.GET("/movies", func(ctx *gin.Context) {
		movies, err := movieService.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, movies)
	})

	apiConnection.GET("/movies/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		movie, err := movieService.GetByID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	})

	apiConnection.POST("/movies", func(ctx *gin.Context) {
		var movie movies.Movie

		if err := ctx.ShouldBindJSON(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := movieService.Create(movie)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	})

	apiConnection.DELETE("movies/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := movieService.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
	})

	// Start API
	apiConnection.Run(":" + PORT)
}
