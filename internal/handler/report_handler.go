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

type ReportHandler struct {
	svc *service.ReportService
}

func NewReportHandler(svc *service.ReportService) *ReportHandler {
	return &ReportHandler{
		svc: svc,
	}
}

func (h *ReportHandler) CreateReport(c *fiber.Ctx) error {
	var req dto.CreateReportRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	projectID, err := uuid.Parse(req.ProjectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID"})
	}

	urlPDF := req.URLPDF

	if urlPDF == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "URLPDF is required"})
	}

	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	report, err := h.svc.CreateReport(ctx, projectID, userID, urlPDF)
	if err != nil {
		return util.WriteError(c, err)
	}
	resp := dto.ReportResponse{
		ID:         report.ID.String(),
		ProjectID:  report.ProjectID.String(),
		URLPDF:     report.URLPDF,
		GeneratedAt: report.GeneratedAt.Format(time.RFC3339),
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *ReportHandler) GetLogByID(c *fiber.Ctx) error {
	reportID := c.Params("id")
	if reportID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "report ID is required"})
	}

	reportIDParsed, err := uuid.Parse(reportID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid report ID format"})
	}
	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	report, err := h.svc.GetReportByID(ctx, reportIDParsed)

	if err != nil {
		return util.WriteError(c, err)
	}
	resp := dto.ReportResponse{
		ID:         report.ID.String(),
		ProjectID:  report.ProjectID.String(),
		URLPDF:     report.URLPDF,
		GeneratedAt: report.GeneratedAt.Format(time.RFC3339),
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *ReportHandler) GetReportsByProject(c *fiber.Ctx) error {
	projectID := c.Params("projectID")

	if projectID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "project ID is required"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	projectIDParsed, err := uuid.Parse(projectID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid project ID format"})
	}


	reports, err := h.svc.GetReportsByProject(ctx, projectIDParsed)
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := make([]dto.ReportResponse, 0, len(reports))
	for _, report := range reports {
		resp = append(resp, dto.ReportResponse{
			ID:         report.ID.String(),
			ProjectID:  report.ProjectID.String(),
			URLPDF:     report.URLPDF,
			GeneratedAt: report.GeneratedAt.Format(time.RFC3339),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}


func (h *ReportHandler) UpdateReport(c *fiber.Ctx) error {
	var req dto.UpdateReportRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	reportID := c.Params("id")
	if reportID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "report ID is required",
		})
	}

	reportIDParsed, err := uuid.Parse(reportID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid report ID format",
		})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	report := &model.Report{
		ID:     reportIDParsed,
		URLPDF: req.URLPDF,
	}

	updatedReport, err := h.svc.UpdateReport(ctx, report)
	if err != nil {
		return util.WriteError(c, err)
	}

	resp := dto.ReportResponse{
		ID:         updatedReport.ID.String(),
		ProjectID:  updatedReport.ProjectID.String(),
		URLPDF:     updatedReport.URLPDF,
		GeneratedAt: updatedReport.GeneratedAt.Format(time.RFC3339),
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *ReportHandler) DeleteReport(c *fiber.Ctx) error {
	reportID := c.Params("id")
	if reportID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "report ID is required"})
	}

	reportIDParsed, err := uuid.Parse(reportID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid report ID format"})
	}

	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	if err := h.svc.DeleteReport(ctx, reportIDParsed); err != nil {
		return util.WriteError(c, err)
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}


