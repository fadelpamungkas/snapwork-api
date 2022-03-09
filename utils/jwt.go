package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("NOTTOKEN")

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return webtoken, err
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
