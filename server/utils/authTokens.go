package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(email string, isRefresh bool) (string, error) {
	var key string
	var duration time.Duration
	if isRefresh {
		key = os.Getenv("REFRESH_TOKEN_KEY")
		duration = RefreshTokenExp
	} else {
		key = os.Getenv("AUTH_TOKEN_KEY")
		duration = AuthTokenExp
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(duration).Unix(),
	})

	token, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJWT(token string, isRefresh bool) (string, error) {
	var key string
	var tokenType string
	if isRefresh {
		key = os.Getenv("REFRESH_TOKEN_KEY")
		tokenType = "refresh"
	} else {
		key = os.Getenv("AUTH_TOKEN_KEY")
		tokenType = "auth"
	}

	tokenJWT, err := jwt.ParseWithClaims(token, &jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", errors.New(tokenType + " token expired")
		}

		return "", err
	}

	claims, ok := tokenJWT.Claims.(*jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to parse jwt claims")
	}

	email, _ := (*claims)["sub"].(string)

	return email, nil
}
