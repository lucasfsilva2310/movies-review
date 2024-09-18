package ratings

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasfsilva2310/movies-review/internal/middlewares"
	"github.com/lucasfsilva2310/movies-review/internal/utils"
)

func RegisterRatingRoutes(apiConnection *gin.Engine, service *RatingService) {
	ratingsURL := apiConnection.Group("/ratings")

	ratingsURL.POST("/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
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

	ratingsURL.GET("/", middlewares.AuthMiddleware(), func(ctx *gin.Context) {

		ratings, errorRequesting := service.GetAll()

		if errorRequesting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorRequesting.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, ratings)
	})

	ratingsURL.GET("/movie/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {

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

	ratingsURL.GET("/user/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {

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

	ratingsURL.PATCH("/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		}

		var rating RatingUpdate

		if err := ctx.ShouldBindJSON(&rating); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Body"})
		}

		username, errorUsername := utils.GetUsernameFromContext(ctx)

		if errorUsername != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errorUsername.Error()})
		}

		err := service.UpdateRating(id, username, rating)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	ratingsURL.DELETE("/:id", middlewares.AdminMiddleware(), func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		username, errorUsername := utils.GetUsernameFromContext(ctx)

		if errorUsername != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errorUsername.Error()})
		}

		err := service.Delete(id, username)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

}
