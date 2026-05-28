package features

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

type Token struct {
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

type TokenBuilder struct {
}

func NewTokenBuilder() *TokenBuilder {
	jwtKey = []byte(os.Getenv("JWT_PASS"))
	return &TokenBuilder{}
}

func (t *TokenBuilder) CreateUserToken(userID, userEmail string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &Token{
		UserID:    userID,
		UserEmail: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *TokenBuilder) ValidateUserToken(tokenString string) (*Token, error) {
	claims := &Token{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}
	return claims, nil
}
