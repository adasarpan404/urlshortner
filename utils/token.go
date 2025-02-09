package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	jwt.RegisteredClaims
}

func generateToken(email string, firstname string, id string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		FirstName: firstname,
		Email:     email,
		ID:        id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(environment.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
