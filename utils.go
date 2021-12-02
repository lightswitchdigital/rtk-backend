package main

import "github.com/golang-jwt/jwt"

func generateJwtToken(access_token string, secret []byte) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": []byte(access_token),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
