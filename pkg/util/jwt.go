package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTProvider interface{
	Generate(userID uuid.UUID) (string, error)
}

type JWTImpl struct {
	secret []byte
	exp time.Duration
}

func NewJWT(secret string, exp time.Duration) *JWTImpl {
	return &JWTImpl{
		secret: []byte(secret),
		exp: exp,
	}
}

func (j *JWTImpl) Generate(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(j.exp).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}