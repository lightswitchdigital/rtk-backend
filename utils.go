package main

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func generateJwtToken(data map[string]interface{}) (string, string, error) {

	claims := jwt.MapClaims{}

	for key, value := range data {
		claims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": 1,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	refTokenString, err := refToken.SignedString(JWT_SECRET_TOKEN)
	if err != nil {
		return "", "", err
	}

	return tokenString, refTokenString, err
}
