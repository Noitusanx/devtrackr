package util

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var ErrInvalidToken = errors.New("invalid or expired token")


type AppError struct {
	Status int

	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func ErrConflict(msg string) error {
	return &AppError{
		Status: http.StatusConflict,
		Message: msg,
	}	
}

func ErrNotFound(msg string) error {
	return &AppError{
		Status: http.StatusNotFound,
		Message: msg,
	}
}

func ErrUnauthorized(msg string) error {
	return &AppError{
		Status: http.StatusUnauthorized,
		Message: msg,
	}
}

func ErrBadRequest(msg string) error {
	return &AppError{
		Status: http.StatusBadRequest,
		Message: msg,
	}
}

func WriteError(c *fiber.Ctx, err error) error {
	if appErr, ok := err.(*AppError); ok {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error": appErr.Message,
		})
	}
	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"error": "internal server error",
	})
}

func ErrInternalServer(msg string) *AppError {
	return &AppError{Status: http.StatusInternalServerError, Message: msg}
}
