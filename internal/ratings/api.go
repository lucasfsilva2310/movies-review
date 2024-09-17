package ratings

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lucasfsilva2310/movies-review/internal/middlewares"
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

		err := service.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

}
