package watchedMovies

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lucasfsilva2310/movies-review/internal/middlewares"
)

func RegisterWatchedMoviesRoutes(apiConnection *gin.Engine, service *WatchedMovieService) {
	watchedMoviesURL := apiConnection.Group("/watched-movies")

	watchedMoviesURL.POST("/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		var watchedMovie WatchedMovie

		if err := ctx.ShouldBindJSON(&watchedMovie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := service.Create(watchedMovie)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	watchedMoviesURL.GET("/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		watchedMovies, err := service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, watchedMovies)
	})

	watchedMoviesURL.GET("/user/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		watchedMovies, err := service.GetAllByUserID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, watchedMovies)
	})

	watchedMoviesURL.GET("/movie/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		watchedMovies, err := service.GetAllByMovieID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, watchedMovies)
	})

	watchedMoviesURL.PATCH("/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		}

		user, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found in token"})
		}

		userMap, ok := user.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token data"})
			return
		}

		username, _ := userMap["username"].(string)

		err := service.UpdateWatchedMovie(id, username)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	watchedMoviesURL.DELETE("/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		err := service.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})
}
