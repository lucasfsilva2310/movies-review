package movies

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(apiConnection *gin.Engine, service *MovieService) {
	moviesURL := apiConnection.Group("/movies")

	moviesURL.GET("/", func(ctx *gin.Context) {
		movies, errorRequesting := service.GetAll()

		if errorRequesting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorRequesting.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, movies)
	})

	moviesURL.GET("/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		movie, errorRequesting := service.GetByID(id)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	})

	moviesURL.POST("/", func(ctx *gin.Context) {
		var movie MovieReturn

		if errorBinding := ctx.ShouldBindJSON(&movie); errorBinding != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errorBinding.Error()})
			return
		}

		errorRequesting := service.Create(movie)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movie)
	})

	moviesURL.DELETE("/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		errorRequesting := service.Delete(id)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
	})
}
