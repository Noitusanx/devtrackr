package util

import (
	"devtracker/internal/domain/dto"
	"os"
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


func GenerateJWT(id string, email string)(string, error){
	secret := os.Getenv("JWT_SECRET");
	if secret == "" {
		secret = "your-super-secret-jwt-key"
	}

	claims := dto.Claims{
		ID: id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "devtracker-api",
			Subject: id,
		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParsingJWT(tokenString string)(*dto.Claims, error){
	secret := os.Getenv("JWT_SECRET");

	if secret == "" {
		secret = "your-super-secret-jwt-key"
	}

	token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
			}
			return []byte(secret), nil
		})

	   if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, ErrInvalidToken
	
}