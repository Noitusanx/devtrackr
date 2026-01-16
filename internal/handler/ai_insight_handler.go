package handler

import (
	"context"
	"strings"
	"time"

	"devtracker/internal/domain/dto"
	"devtracker/internal/domain/model"
	"devtracker/internal/service"
	"devtracker/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AIInsightHandler struct {
	svc *service.AIInsightService
}


func NewAIInsightHandler(svc *service.AIInsightService) *AIInsightHandler {
	return &AIInsightHandler{
		svc: svc,
	}
}


func (h *AIInsightHandler) CreateInsight(c *fiber.Ctx) error {
	var req dto.AIInsightRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	projectID, err := uuid.Parse(req.ProjectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	insightType := model.InsightType(strings.ToLower(req.InsightType))
	if insightType != model.InsightSummary && insightType != model.InsightProgress {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid insight type"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	insight, err := h.svc.CreateInsight(ctx, projectID, insightType, req.Content)
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := dto.AIInsightResponse{
		ID:          insight.ID.String(),
		ProjectID:   insight.ProjectID.String(),
		InsightType: string(insight.Type),
		Content:     insight.Content,
		GeneratedAt: insight.GeneratedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AIInsightHandler) GetInsightByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "insight ID is required"})
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid insight ID"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	insight, err := h.svc.GetInsightByID(ctx, id)
	if err != nil {
		return util.WriteError(c, err)
	}

	if insight == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "insight not found"})
	}

	resp := dto.AIInsightResponse{
		ID:          insight.ID.String(),
		ProjectID:   insight.ProjectID.String(),
		InsightType: string(insight.Type),
		Content:     insight.Content,
		GeneratedAt: insight.GeneratedAt,
	}

	return c.JSON(resp)
}

func (h *AIInsightHandler) GetInsightsByProject(c *fiber.Ctx) error {
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

	insights, err := h.svc.GetInsightsByProject(ctx, projectID)
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := make([]dto.AIInsightResponse, len(insights))
	for i, insight := range insights {
		resp[i] = dto.AIInsightResponse{
			ID:          insight.ID.String(),
			ProjectID:   insight.ProjectID.String(),
			InsightType: string(insight.Type),
			Content:     insight.Content,
			GeneratedAt: insight.GeneratedAt,
	
		}
	}

	return c.JSON(resp)
}

func (h *AIInsightHandler) UpdateInsight(c *fiber.Ctx) error {
	var req dto.AIInsightUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	id, err := uuid.Parse(*req.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid insight ID"})
	}

	projectID, err := uuid.Parse(*req.ProjectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	insightType := model.InsightType(strings.ToLower(*req.InsightType))
	if insightType != model.InsightSummary && insightType != model.InsightProgress {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid insight type"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	insight, err := h.svc.UpdateInsight(ctx, &model.AIInsight{
		ID:          id,
		ProjectID:   projectID,
		Type:        insightType,
		Content:     *req.Content,
	})
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := dto.AIInsightResponse{
		ID:          insight.ID.String(),
		ProjectID:   insight.ProjectID.String(),
		InsightType: string(insight.Type),
		Content:     insight.Content,
		GeneratedAt: insight.GeneratedAt,
	}

	return c.JSON(resp)
}


func (h *AIInsightHandler) GetInsightByProjectAndTypeAndDate(c *fiber.Ctx) error {
	projectIDStr := c.Params("projectID")
	if projectIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "project ID is required"})
	}

	projectID, err := uuid.Parse(projectIDStr)
	if err != nil || projectID == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	insightTypeStr := c.Params("insightType")
	if insightTypeStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "insight type is required"})
	}
	insightType := model.InsightType(strings.ToLower(insightTypeStr))
	if insightType != model.InsightSummary && insightType != model.InsightProgress {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid insight type"})
	}

	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "date is required"})
	}
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid date format"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	insight, err := h.svc.GetInsightByProjectAndTypeAndDate(ctx, projectID, insightType, date)
	if err != nil {
		return util.WriteError(c, err)
	}

	if insight == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "insight not found"})
	}

	resp := dto.AIInsightResponse{
		ID:          insight.ID.String(),
		ProjectID:   insight.ProjectID.String(),
		InsightType: string(insight.Type),
		Content:     insight.Content,
		GeneratedAt: insight.GeneratedAt,
	}

	return c.JSON(resp)
}

func (h *AIInsightHandler) DeleteInsight(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "insight ID is required"})
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid insight ID"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	err = h.svc.DeleteInsight(ctx, id)
	if err != nil {
		return util.WriteError(c, err)
	}

	return c.Status(fiber.StatusNoContent).SendString("")
}