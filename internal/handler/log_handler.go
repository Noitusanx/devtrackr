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

type logHandler struct {
	svc *service.LogService
}

func NewLogHandler(svc *service.LogService) *logHandler {
	return &logHandler{
		svc: svc,
	}
}

func (h *logHandler) CreateLog(c *fiber.Ctx) error {
	var req dto.CreateLogRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	projectID, err := uuid.Parse(req.ProjectID)
	if err != nil || projectID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(uuid.UUID); 
	
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	loggedAt, err := time.Parse(time.RFC3339, req.LoggedAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid logged at format"})
	}


	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.CreateLog(ctx, projectID, userID, req.Text, req.DurationMinutes, loggedAt)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "log created successfully"})
}

func (h *logHandler) GetLogByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "log ID is required"})
	}

	logID, err := uuid.Parse(idStr)
	if err != nil || logID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid log ID"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	log, err := h.svc.GetLogByID(ctx, logID)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.JSON(log)
}

func (h *logHandler) GetLogsByUser(c *fiber.Ctx) error {
	userIDVal := c.Locals("userID")

	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}


	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	logs, err := h.svc.GetLogsByUser(ctx, userID)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.JSON(logs)
}

func (h *logHandler) GetLogsByProject(c *fiber.Ctx) error {
	projectIDStr := c.Params("projectID")
	if projectIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "project ID is required"})
	}

	projectID, err := uuid.Parse(projectIDStr)
	if err != nil || projectID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	logs, err := h.svc.GetLogsByProject(ctx, projectID)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.JSON(logs)
}

func (h *logHandler) UpdateLog(c *fiber.Ctx) error {
	var req dto.UpdateLogRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	logIDStr := c.Params("id")
	if logIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "log ID is required"})
	}

	logID, err := uuid.Parse(logIDStr)
	if err != nil || logID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid log ID"})
	}

	projectID, err := uuid.Parse(*req.ProjectID)
	if err != nil || projectID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(uuid.UUID); 
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	log, err := h.svc.GetLogByID(c.Context(), logID)
	if err != nil {
		return util.WriteError(c, err)
	}

	if req.Text != nil {
		log.Text = *req.Text
	}

	if req.DurationMinutes != nil {
		log.DurationMinutes = *req.DurationMinutes
	}

	if req.LoggedAt != nil {
		if *req.LoggedAt == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "logged at cannot be empty"})
		}
		t, err := time.Parse(time.RFC3339, *req.LoggedAt)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid logged at format"})
		}
		log.LoggedAt = t
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.UpdateLog(ctx, userID, log)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
