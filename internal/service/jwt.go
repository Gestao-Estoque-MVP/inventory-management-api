package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/diogoX451/inventory-management-api/internal/graph/model"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaim struct {
	ID   string `json:"id"`
	Role model.Role
	jwt.RegisteredClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}


func JwtGenerate(_ context.Context, userID string, role model.Role) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID:   userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
	})

	token, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
