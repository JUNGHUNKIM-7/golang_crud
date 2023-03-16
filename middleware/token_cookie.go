package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, td time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       email,
		"issuer":     "localhost",
		"created_at": time.Now().Unix(),
		"exp":        time.Now().Add(td).Unix(),
	})

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return "", errors.New("failed gen token")
	}
	return tokenStr, nil
}

func VerifyToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad signed method received")
		}
		return SECRET, nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}
