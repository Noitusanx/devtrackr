package middleware

import (
	"devtracker/internal/domain/dto"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func JWTProtected() fiber.Handler{
	return func(c *fiber.Ctx) error{
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid Authorization",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer"{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid Authorization, expected: Bearer <token>",
			})
		}

		tokenString := tokenParts[1];

		token, err := jwt.ParseWithClaims(tokenString, &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				secret = "your-super-secret-jwt-key" 
			}
			return []byte(secret), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
			uid, err := uuid.Parse(claims.ID)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid user id",
				})
			}



			c.Locals("userID", uid)
			c.Locals("userEmail", claims.Email)
		
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})

	}
}

func OptionalJWT() fiber.Handler {
	return func (c *fiber.Ctx) error {
		authHeader := c.Get("Authorization");

		if authHeader == "" {
			c.Next()
		}

		tokenParts:= strings.Split(authHeader, " ")

		if len(tokenParts) == 2 || tokenParts[0] == "Bearer" {
			tokenString := tokenParts[1];

			token, err := jwt.ParseWithClaims(tokenString, dto.Claims{}, func(t *jwt.Token) (interface{}, error) {
				secret := os.Getenv("JWT_SECRET")
				if secret == "" {
					secret = "your-super-secret-jwt-key"
				}
				return []byte(secret), nil
			})

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid token or expired token",
				})
			}

			if claims, ok := token.Claims.(*dto.Claims); ok && token.Valid {
				c.Locals("userID", claims.ID)
				c.Locals("email", claims.Email)

			}	
		}

		return c.Next()
	}

}