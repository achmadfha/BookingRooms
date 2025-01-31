package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"strconv"
	"time"
)

func GenerateToken(id uuid.UUID, position string) (tokenString string, err error) {
	expired := os.Getenv("TOKEN_EXPIRED")
	secret := os.Getenv("SECRET_TOKEN")
	exp, err := strconv.Atoi(expired)
	if err != nil {
		return "", err
	}

	expiredTime := time.Now().Add(time.Duration(exp) * time.Hour)

	claims := jwt.MapClaims{
		"id":      id,
		"role":    position,
		"expired": expiredTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
