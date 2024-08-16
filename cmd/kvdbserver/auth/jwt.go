package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTOptions struct {
	// The signing method
	SignMethod jwt.SigningMethod
	// The signing key
	SignKey string
	// Time to live in seconds
	TTL time.Duration
}

// GenerateJWT generates a new JWT.
func GenerateJWT(opts JWTOptions, username string) (string, error) {
	token := jwt.NewWithClaims(opts.SignMethod, jwt.MapClaims{
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(opts.TTL).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(opts.SignKey))
	if err != nil {
		return "", err
	}

	return tokenStr, err
}
