package movieComments

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterMovieCommentsRoutes(apiConnection *gin.Engine, service *MovieCommentService) {
	movieCommentsURL := apiConnection.Group("/movie-comments")

	movieCommentsURL.POST("/", func(ctx *gin.Context) {
		var movieComment MovieComment

		if err := ctx.ShouldBindJSON(&movieComment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := service.Create(movieComment)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	movieCommentsURL.GET("/", func(ctx *gin.Context) {
		movieComments, err := service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movieComments)
	})

	movieCommentsURL.GET("/user/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		movieComments, err := service.GetAllByUserID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movieComments)
	})

	movieCommentsURL.GET("/movie/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		movieComments, err := service.GetAllByMovieID(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, movieComments)
	})
}
