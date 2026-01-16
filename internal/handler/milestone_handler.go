package handler

import (
	"context"
	"time"

	"devtracker/internal/domain/dto"
	"devtracker/internal/domain/model"
	"devtracker/internal/service"
	"devtracker/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MilestoneHandler struct {
	svc *service.MilestoneService
}


func NewMilestoneHandler(svc *service.MilestoneService) *MilestoneHandler {
	return &MilestoneHandler{
		svc: svc,
	}
}


func (h *MilestoneHandler) CreateMilestone(c *fiber.Ctx) error {
	var req dto.CreateMilestoneRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	projectID := c.Params("id")
	if projectID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "project ID is required"})
	}

	var dueDatePtr *time.Time
	if req.DueDate != "" {
		dueDate, err := time.Parse(time.RFC3339, req.DueDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid due date format"})
		}
		dueDatePtr = &dueDate
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	
	projectIDParsed, err := uuid.Parse(projectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID format"})
	}

	if req.Status == "" {
		req.Status = model.StatusPending
	}


	milestone, err := h.svc.CreateMilestone(ctx, projectIDParsed, req.Name, req.OrderIdx, req.Status, dueDatePtr)
	if err != nil {
		return util.WriteError(c, err)
	}



	resp := dto.MilestoneResponse{
		ID:          milestone.ID.String(),
		ProjectID: projectID,
		Name:        milestone.Name,
		OrderIdx:    milestone.OrderIdx,
		Status:      milestone.Status,
		DueDate:     util.FormatPtr(milestone.DueDate),
		CompletedAt: util.FormatPtr(milestone.CompletedAt),
		CreatedAt:   milestone.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   milestone.UpdatedAt.Format(time.RFC3339),
	}
	return c.Status(fiber.StatusCreated).JSON(resp)

}

func(h *MilestoneHandler) GetMilestonesByProject(c *fiber.Ctx) error {
	projectID := c.Params("id")
	if projectID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "project ID is required"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	UserID := c.Locals("userID").(uuid.UUID)

	projectIDParsed, err := uuid.Parse(projectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID format"})
	}

	milestones, err := h.svc.GetMilestonesByProject(ctx, UserID, projectIDParsed)
	if err != nil {
		return util.WriteError(c, err)
	}

	var resp []dto.MilestoneResponse
	for _, m := range milestones {
		resp = append(resp, dto.MilestoneResponse{
			ID:          m.ID.String(),
			ProjectID: projectID,
			Name:        m.Name,
			OrderIdx:    m.OrderIdx,
			Status:      m.Status,
			DueDate:     util.FormatPtr(m.DueDate),
			CompletedAt: util.FormatPtr(m.CompletedAt),
			CreatedAt:   m.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   m.UpdatedAt.Format(time.RFC3339),
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *MilestoneHandler) GetMilestoneByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "milestone ID is required"})
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid milestone ID format"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	milestone, err := h.svc.GetMilestoneByID(ctx, id)
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := dto.MilestoneResponse{
		ID:          milestone.ID.String(),
		ProjectID: milestone.ProjectID.String(),
		Name:        milestone.Name,
		OrderIdx:    milestone.OrderIdx,
		Status:      milestone.Status,
		DueDate:     util.FormatPtr(milestone.DueDate),
		CompletedAt: util.FormatPtr(milestone.CompletedAt),
		CreatedAt:   milestone.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   milestone.UpdatedAt.Format(time.RFC3339),
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *MilestoneHandler) UpdateMilestone(c *fiber.Ctx) error {
	var req dto.UpdateMilestoneRequest

	

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "milestone ID is required"})
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid milestone ID format"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	milestone, err := h.svc.GetMilestoneByID(ctx, id)
	if err != nil {
		return util.WriteError(c, err)
	}

	if req.Name != "" {
		milestone.Name = req.Name
	}

	if req.OrderIdx != nil {
		milestone.OrderIdx = *req.OrderIdx
	}

	if req.Status != "" {
		switch req.Status{
			case "pending", "in_progress", "done":
				milestone.Status = model.MilestoneStatus(req.Status)
			default:
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid status"})
		}
	}


	if req.DueDate != nil { 
	if *req.DueDate == "" {        
		milestone.DueDate = nil
	} else {
		d, err := time.Parse(time.RFC3339, *req.DueDate)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid due date format"})
		}
		milestone.DueDate = &d
	}
}

	if err := h.svc.UpdateMilestone(ctx, milestone); err != nil {
		return util.WriteError(c, err)
	}

	if milestone.Status == model.StatusDone && milestone.CompletedAt == nil {
		now := time.Now()
		milestone.CompletedAt = &now
	}
	if milestone.Status != model.StatusDone {
		milestone.CompletedAt = nil
}

	resp := dto.MilestoneResponse{
		ID:          milestone.ID.String(),
		ProjectID: milestone.ProjectID.String(),
		Name:        milestone.Name,
		OrderIdx:    milestone.OrderIdx,
		Status:      milestone.Status,
		DueDate:     util.FormatPtr(milestone.DueDate),
		CompletedAt: util.FormatPtr(milestone.CompletedAt),
		CreatedAt:   milestone.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   milestone.UpdatedAt.Format(time.RFC3339),
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}


func (h *MilestoneHandler) DeleteMilestone(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "milestone ID is required"})
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid milestone ID format"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	if err := h.svc.DeleteMilestone(ctx, id); err != nil {
		return util.WriteError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

