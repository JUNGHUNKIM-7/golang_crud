package middleware

import (
	"errors"
	"net/http"
	"runner/src/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func JwtTokenCheckFromBearer(c *gin.Context) {
	jwtToken, err := ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := VerifyToken(jwtToken)
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
