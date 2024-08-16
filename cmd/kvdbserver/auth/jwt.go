package auth

import (
	"context"
	"errors"
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

type AuthInfo struct {
	Username string
}

// GenerateJWT generates a new JWT token.
func GenerateJWT(ctx context.Context, opts *JWTOptions, username string) (string, error) {
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

// ValidateJWT validates JWT token.
func ValidateJWT(ctx context.Context, tokenStr string, opts *JWTOptions) (*AuthInfo, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != opts.SignMethod.Alg() {
			return nil, errors.New("invalid signing method")
		}
		return opts.SignKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, errors.New("failed to get claims from JWT token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("failed to get username from claims")
	}

	return &AuthInfo{Username: username}, nil
}
