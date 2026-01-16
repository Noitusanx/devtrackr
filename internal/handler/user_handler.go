package handler

import (
	"context"
	"time"

	"devtracker/internal/domain/dto"
	"devtracker/internal/service"
	"devtracker/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Validation (bisa pakai validator library)
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email, password, and name are required",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// Service melakukan SEMUA business logic:
	// - Check duplicate email
	// - Hash password
	// - Create user
	// - Generate JWT token
	user, token, err := h.svc.Register(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return util.WriteError(c, err)
	}

	// Handler hanya format response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": token,
	})
}


// Login
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Validation
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// Service melakukan SEMUA business logic:
	// - Find user
	// - Validate password
	// - Generate JWT token
	user, token, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return util.WriteError(c, err)
	}

	// Handler hanya format response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
		"token": token,
	})
}


func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	userID := c.Locals("userID").(uuid.UUID)

	user, err := h.svc.GetCurrentUser(ctx, userID)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.Status(200).JSON(dto.User{
		ID:    userID.String(),
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
    var req dto.UpdateUserRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "invalid request body",
        })
    }

    userID := c.Locals("userID").(uuid.UUID)



    ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
    defer cancel()

    // Call service with only necessary data
    user, err := h.svc.UpdateProfile(ctx, userID, req.Name)
    if err != nil {
        return util.WriteError(c, err)
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Profile updated successfully",
        "user": fiber.Map{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        },
    })
}

func (h *UserHandler) DeleteAccount(c *fiber.Ctx) error {

	userID := c.Locals("userID").(uuid.UUID)

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second);

	defer cancel()

	err := h.svc.DeleteAccount(ctx, userID)

	if err != nil{
		return util.WriteError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

