package dto

import "github.com/golang-jwt/jwt/v5"

type RegisterRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User User `json:"user"`
}


type Claims struct {
	ID string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
