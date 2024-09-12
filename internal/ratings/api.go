package ratings

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRatingRoutes(apiConnection *gin.Engine, service *RatingService) {
	ratingsURL := apiConnection.Group("/ratings")

	ratingsURL.POST("/", func(ctx *gin.Context) {
		var rating Rating

		if err := ctx.ShouldBindJSON(&rating); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.Create(rating)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	ratingsURL.GET("/", func(ctx *gin.Context) {

		ratings, errorRequesting := service.GetAll()

		if errorRequesting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorRequesting.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, ratings)
	})

	ratingsURL.GET("/movie/:id", func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		ratings, errorRequesting := service.GetAllByMovieID(id)

		if errorRequesting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorRequesting.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, ratings)
	})

	ratingsURL.GET("/user/:id", func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		ratings, errorRequesting := service.GetAllByUserID(id)

		if errorRequesting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorRequesting.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, ratings)
	})

	ratingsURL.DELETE("/:id", func(ctx *gin.Context) {

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
