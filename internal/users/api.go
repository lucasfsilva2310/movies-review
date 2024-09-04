package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(apiConnection *gin.Engine, service *UserService) {
	usersURL := apiConnection.Group("/users")

	usersURL.POST("/create", func(ctx *gin.Context) {
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

		ctx.JSON(http.StatusOK, user)
	})
}
