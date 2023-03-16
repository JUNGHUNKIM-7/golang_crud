package middleware

import (
	"errors"
	"net/http"
	"runner/src/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// todo .env
var SECRET = []byte("supersecret")

func GetEmailFromCookie(c *gin.Context, tokenName string) (string, error) {
	cookie, err := c.Cookie("at")
	if err != nil {
		return "", errors.New("failed to get cookie")
	}

	token, err := VerifyToken(cookie)
	if err != nil {
		return "", errors.New("failed to verify token")
	}

	if userEmail, ok := token.Claims.(jwt.MapClaims)["user"].(string); ok {
		return userEmail, nil
	}
	return "", errors.New("failed to get token body")
}

func CheckJwtFromCookie(c *gin.Context) {
	cookie, err := c.Cookie("at")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := VerifyToken(cookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}
