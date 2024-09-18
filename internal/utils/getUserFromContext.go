package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUsernameFromContext(c *gin.Context) (string, error) {
	user, exists := c.Get("user")
	if !exists {
		return "", errors.New("user not found in context")
	}

	userMap, ok := user.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token data")
	}

	username, _ := userMap["username"].(string)

	return username, nil
}
