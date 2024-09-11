package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasfsilva2310/movies-review/internal/users"
)

func RegisterAuthRoutes(apiConnection *gin.Engine, service *AuthService) {
	authURL := apiConnection.Group("/auth")

	authURL.POST("/login", func(ctx *gin.Context) {
		var loginRequest LoginRequest

		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := service.Login(loginRequest)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"token": token})
	})

	authURL.POST("/register", func(ctx *gin.Context) {
		var user users.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.Register(user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})
}
