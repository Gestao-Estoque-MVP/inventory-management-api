package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWT() *JWT {
	privateBytes, err := os.ReadFile("private_key.pem")
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		panic(err)
	}

	publicBytes, err := os.ReadFile("public_key.pem")
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		panic(err)
	}

	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j *JWT) GenerateToken(id int32, role string) (string, error) {
	now := time.Now().UTC().Local()

	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = now.Add(time.Hour * 24).Unix()
	claims["iat"] = now.Unix()
	claims["iss"] = "math_api"
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(j.privateKey)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *JWT) ValidateToken(token string) (int32, string, error) {

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})

	if err != nil {
		return 0, "", err
	}

	id, ok := claims["id"].(int32)
	if !ok {
		return 0, "", fmt.Errorf("id not found")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", fmt.Errorf("role not found")
	}

	return id, role, nil

}
