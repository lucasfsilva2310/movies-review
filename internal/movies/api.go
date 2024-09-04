package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(apiConnection *gin.Engine, service *MovieService) {
	moviesURL := apiConnection.Group("/movies")

	moviesURL.GET("/", func(ctx *gin.Context) {
		movies, err := service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, movies)
	})

	moviesURL.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		movie, err := service.GetByID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	})

	moviesURL.POST("/", func(ctx *gin.Context) {
		var movie Movie

		if err := ctx.ShouldBindJSON(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.Create(movie)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	})

	moviesURL.DELETE("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := service.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
	})
}
