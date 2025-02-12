package utils

import (
	"time"

	"github.com/adasarpan404/urlshortner/environment_variables"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(email string, firstname string, id string) (string, error) {
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

	tokenString, err := token.SignedString([]byte(environment_variables.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(signedToken string) (claim *Claims, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(t *jwt.Token) (interface{}, error) { return []byte(environment.SECRET_KEY), nil })
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		msg = err.Error()
		return
	}
	return claims, msg
}
