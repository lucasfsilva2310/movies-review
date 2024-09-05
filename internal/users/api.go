package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(apiConnection *gin.Engine, service *UserService) {
	usersURL := apiConnection.Group("/users")

	usersURL.POST("/", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := service.Create(user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, nil)
	})

	usersURL.GET("/:id", func(ctx *gin.Context) {
		var user UserReturn

		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		user, errorRequesting := service.GetById(id)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	usersURL.GET("/username/:username", func(ctx *gin.Context) {
		var user UserReturn

		username := ctx.Param("username")

		user, errorRequesting := service.GetByUsername(username)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)

	})

	usersURL.DELETE("/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, errorConverting := strconv.Atoi(idParam)

		if errorConverting != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		errorRequesting := service.DeleteById(id)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

	usersURL.DELETE("/username/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")

		errorRequesting := service.DeleteByUsername(username)

		if errorRequesting != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errorRequesting.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})
}
