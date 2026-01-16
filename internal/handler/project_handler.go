package handler

import (
	"context"
	"log"
	"time"

	"devtracker/internal/domain/dto"
	"devtracker/internal/domain/model"
	"devtracker/internal/service"
	"devtracker/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


	type ProjectHandler struct {
		svc *service.ProjectService
	}

	func NewProjectHandler(svc *service.ProjectService) *ProjectHandler {
		return &ProjectHandler{
			svc: svc,
		}
	}

	func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
		var req dto.CreateProjectRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
		}

		userID := c.Locals("userID").(uuid.UUID)

		var deadlinePtr *time.Time
		if req.Deadline != "" {
			d, err := time.Parse(time.RFC3339, req.Deadline)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "invalid deadline format"})
			}
			deadlinePtr = &d
		}

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		project, err := h.svc.CreateProject(ctx, userID, req.Name, deadlinePtr)
		if err != nil {
			return util.WriteError(c, err)
		}

		resp := dto.ProjectResponse{
			ID:        project.ID.String(),
			Name:      project.Name,
			Deadline:  util.FormatPtr(project.Deadline),
			CreatedAt: project.CreatedAt.Format(time.RFC3339),
			UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		}
		return c.Status(201).JSON(resp)
	}




	func (h *ProjectHandler) GetProjectByID(c *fiber.Ctx) error {

		idStr := c.Params("id")
		log.Printf("DEBUG: Received ID param: '%s' (len=%d)", idStr, len(idStr))


		if idStr == "" {
			return c.Status(400).JSON(fiber.Map{"error": "project ID is required"})
		}

		
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid project ID format"})
		}
		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		project, err := h.svc.GetProjectByID(ctx, id)
		if err != nil {
			return util.WriteError(c, err)
		}

		resp := (dto.ProjectResponse{
			ID:        project.ID.String(),
			Name:      project.Name,
			Deadline:  util.FormatPtr(project.Deadline),
			CreatedAt: project.CreatedAt.Format(time.RFC3339),
			UpdatedAt: project.UpdatedAt.Format(time.RFC3339),
		})

		return c.Status(200).JSON(resp)
	}


	func (h *ProjectHandler) ListProjectsByUser(c *fiber.Ctx) error {
		userID := c.Locals("userID").(uuid.UUID)

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		projects, err := h.svc.ListProjectByUser(ctx, userID)
		if err != nil {
			return util.WriteError(c, err)
		}

		var resp []dto.ProjectResponse
		for _, p := range projects {
			resp = append(resp, dto.ProjectResponse{
				ID:        p.ID.String(),
				Name:      p.Name,
				Deadline:  util.FormatPtr(p.Deadline),
				CreatedAt: p.CreatedAt.Format(time.RFC3339),
				UpdatedAt: p.UpdatedAt.Format(time.RFC3339),
			})
		}

		return c.Status(200).JSON(resp)
	}

	func (h *ProjectHandler) UpdateProject(c *fiber.Ctx) error {
		var req dto.UpdateProjectRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
		}

		idStr := c.Params("id")

		userID := c.Locals("userID").(uuid.UUID)

		if idStr == "" {
			return c.Status(400).JSON(fiber.Map{"error": "project ID is required"})
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid project ID format"})
		}

		var deadlinePtr *time.Time

		if req.Deadline !=""{
			d, err := time.Parse(time.RFC3339, req.Deadline)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{"error": "invalid deadline format"})
			}
			deadlinePtr = &d
		}

		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		project := &model.Project{
			ID:       id,
			UserID:   userID,
			Name:     req.Name,
			Deadline: deadlinePtr,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := h.svc.UpdateProject(ctx, userID, project); err != nil {
			return util.WriteError(c, err)
		}

		return c.JSON(project)
	}




	func (h *ProjectHandler) DeleteProject(c *fiber.Ctx) error {
		idStr :=c.Params("id")



		if idStr == "" {
			return c.Status(400).JSON(fiber.Map{"error": "project ID is required"})
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid project ID format"})
		}

		userID := c.Locals("userID").(uuid.UUID)


		ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
		defer cancel()

		if err := h.svc.DeleteProject(ctx, userID, id); err != nil {
			return util.WriteError(c, err)
		}

		return c.SendStatus(204)

		
	}

	func (h *ProjectHandler) GetProjectLogs(c *fiber.Ctx){
		
	}



		

