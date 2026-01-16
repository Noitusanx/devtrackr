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

type LogHandler struct {
	svc *service.LogService
}

func NewLogHandler(svc *service.LogService) *LogHandler {
	return &LogHandler{
		svc: svc,
	}
}

func (h *LogHandler) CreateLog(c *fiber.Ctx) error {
	var req dto.CreateLogRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	var projectID uuid.UUID
    var err error

 
    projectIDStr := c.Params("id")
    if projectIDStr != "" {
        projectID, err = uuid.Parse(projectIDStr)
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID from URL"})
        }
    } else {
        projectID, err = uuid.Parse(req.ProjectID)
        if err != nil || projectID == uuid.Nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID in body"})
        }
    }
	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(uuid.UUID); 
	
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}


	var milestoneID *uuid.UUID;

	if req.MilestoneID != nil && *req.MilestoneID != ""{
		parsedMilestoneID, err := uuid.Parse(*req.MilestoneID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid milestone ID format"})
		}
		milestoneID = &parsedMilestoneID;
	}

	loggedAt, err := time.Parse(time.RFC3339, req.LoggedAt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid logged at format"})
	}


	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	log, err := h.svc.CreateLog(ctx, projectID, userID, milestoneID, req.Description, req.DurationMinutes, loggedAt)
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := dto.LogResponse{
    ID:              log.ID.String(),
    ProjectID:       log.ProjectID.String(),
    MilestoneID:     util.UUIDPtrToStringPtr(log.MilestoneID),
    UserID:          log.UserID.String(),
    Description:     log.Description,
    DurationMinutes: log.DurationMinutes,
    LoggedAt:        log.LoggedAt.Format(time.RFC3339),
    CreatedAt:       log.CreatedAt.Format(time.RFC3339),
}
	

	return c.Status(201).JSON(resp)
}

func (h *LogHandler) GetLogByID(c *fiber.Ctx) error {
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

func (h *LogHandler) GetLogsByUser(c *fiber.Ctx) error {
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

func (h *LogHandler) GetLogsByProject(c *fiber.Ctx) error {
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


	var response []dto.LogResponse;

	for _, log := range logs {
		var milestoneIDstr *string
		if log.MilestoneID != nil{
			s := log.MilestoneID.String()
			milestoneIDstr = &s
		}

		response = append(response, dto.LogResponse{
			ID: log.ID.String(),
			ProjectID: log.ProjectID.String(),
			MilestoneID: milestoneIDstr,
			UserID: log.UserID.String(),
			Description: log.Description,
			DurationMinutes: log.DurationMinutes,
			LoggedAt: log.LoggedAt.Format(time.RFC3339),
			CreatedAt: log.CreatedAt.Format(time.RFC3339),
		})
	}

	if response == nil {
        response = []dto.LogResponse{}
    }
	
	return c.JSON(response)
}

func (h *LogHandler) UpdateLog(c *fiber.Ctx) error {
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

	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(uuid.UUID); 
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	log, err := h.svc.GetLogByID(c.Context(), logID)
	if err != nil {
		return util.WriteError(c, err)
	}



	if req.Description != nil {
		log.Description = *req.Description
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

	if req.MilestoneID != nil {
		if *req.MilestoneID == ""{
			log.MilestoneID = nil
		} else {
			parsedMilestoneID, err := uuid.Parse(*req.MilestoneID)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid milestone ID"});
			}

			log.MilestoneID = &parsedMilestoneID;
		}
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.UpdateLog(ctx, userID, log)
	if err != nil {
		return util.WriteError(c, err)
	}

	var milestoneIDstr *string

	if log.MilestoneID != nil{
		s := log.MilestoneID.String()
		milestoneIDstr = &s;
	}

	resp := dto.LogResponse{
		ID: log.ID.String(),
		ProjectID: log.ProjectID.String(),
		MilestoneID: milestoneIDstr,
		UserID: log.UserID.String(),
		Description: log.Description,
		DurationMinutes: log.DurationMinutes,
		LoggedAt: log.LoggedAt.Format(time.RFC3339),
		CreatedAt: log.CreatedAt.Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	return c.JSON(resp)
}

func (h *LogHandler) DeleteLog(c *fiber.Ctx) error {
	logIDStr := c.Params("id")
	if logIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "log ID is required"})
	}

	logID, err := uuid.Parse(logIDStr)
	if err != nil || logID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid log ID"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.DeleteLog(ctx, logID)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
