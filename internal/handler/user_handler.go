package handler

import (
	"context"
	"time"

	"devtracker/internal/domain/dto"
	"devtracker/internal/service"
	"devtracker/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler{
	return &UserHandler{
		svc: svc,
	}
}


// Register
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	user, token, err := h.svc.Register(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.Status(201).JSON(dto.AuthResponse{
		Token: token,
		User:  dto.FromModelUser(user),
	})
}


// Login
func (h *UserHandler) Login(c *fiber.Ctx) error{
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	user, token, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return util.WriteError(c, err)
	}
	return c.Status(200).JSON(dto.AuthResponse{
		Token: token,
		User:  dto.FromModelUser(user),
	})
}