package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("KtrYPBh3sNfdEm4UhrnLWz8qCCxcq0Pm1G2D4spgWUaWSxzTrVqXcUpIVvSWbEoY")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
func GenerateToken(name string) (string, error) {
	claims := jwt.MapClaims {
		"UserName": name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(JWT_KEY)
}