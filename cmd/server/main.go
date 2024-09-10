package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Postgres Driver

	movieComments "github.com/lucasfsilva2310/movies-review/internal/MovieComments"
	Configuration "github.com/lucasfsilva2310/movies-review/internal/config"
	"github.com/lucasfsilva2310/movies-review/internal/movies"
	"github.com/lucasfsilva2310/movies-review/internal/ratings"
	"github.com/lucasfsilva2310/movies-review/internal/users"
	"github.com/lucasfsilva2310/movies-review/internal/watchedMovies"
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
	apiConnection := Configuration.CreateApiConnection()

	// Repositories
	movieRepo := movies.NewMovieRepository(Configuration.NewRepository(dbConn))
	userRepo := users.NewUserRepository(Configuration.NewRepository(dbConn))
	ratingRepo := ratings.NewRatingRepository(Configuration.NewRepository(dbConn))
	watchedMoviesRepo := watchedMovies.NewWatchedMovieRepository(Configuration.NewRepository(dbConn))
	movieCommentsRepo := movieComments.NewMovieCommentRepository(Configuration.NewRepository(dbConn))

	// Services
	movieService := movies.NewMovieService(movieRepo)
	userService := users.NewUserService(userRepo)
	ratingService := ratings.NewRatingService(ratingRepo)
	watchedMoviesService := watchedMovies.NewWatchedMovieService(watchedMoviesRepo)
	movieCommentsService := movieComments.NewMovieCommentService(movieCommentsRepo)

	// Endpoints
	movies.RegisterMovieRoutes(apiConnection, movieService)
	users.RegisterUserRoutes(apiConnection, userService)
	ratings.RegisterRatingRoutes(apiConnection, ratingService)
	watchedMovies.RegisterWatchedMoviesRoutes(apiConnection, watchedMoviesService)
	movieComments.RegisterMovieCommentsRoutes(apiConnection, movieCommentsService)

	// Health
	apiConnection.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// Start API
	apiConnection.Run(":" + PORT)
}
