package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: add loadConfig from database/sql

	apiConnection := gin.Default()

	// TODO: Add endpoints List Here
	// Example
	apiConnection.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	apiConnection.Run("0.0.0.0:8080")
}
