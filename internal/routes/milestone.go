package routes

import (
	"devtracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func setupMilestoneRoutes(app fiber.Router, handler *handler.MilestoneHandler) {

	milestones := app.Group("/milestones")

	milestones.Get("/:id", handler.GetMilestoneByID)
	milestones.Put("/:id", handler.UpdateMilestone)
	milestones.Delete("/:id", handler.DeleteMilestone)


	// milestones.Patch("/:id/complete", handler.CompleteMilestone) // TODO
}