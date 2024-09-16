package middlewares

import (
	"fmt"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adminKey := os.Getenv("ADMIN_KEY")

		token := ctx.GetHeader("Admin")

		fmt.Println("Token: ", token)
		fmt.Println("adminKey: ", adminKey)

		if adminKey != token {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
