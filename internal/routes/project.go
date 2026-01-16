package routes

import (
	"devtracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func setupProjectRoutes(
	app fiber.Router,
	projectHandler *handler.ProjectHandler,
	milestoneHandler *handler.MilestoneHandler,
	logHandler *handler.LogHandler,
	reportHandler *handler.ReportHandler,
) {
	projects := app.Group("/projects")

	// Project CRUD
	projects.Post("/", projectHandler.CreateProject)
	projects.Get("/", projectHandler.ListProjectsByUser)
	projects.Get("/:id", projectHandler.GetProjectByID)
	projects.Put("/:id", projectHandler.UpdateProject)
	projects.Delete("/:id", projectHandler.DeleteProject)

	// ✅ Nested: Milestones under Project
	projects.Post("/:id/milestones", milestoneHandler.CreateMilestone)
	projects.Get("/:id/milestones", milestoneHandler.GetMilestonesByProject)

	// ✅ Nested: Logs under Project
	projects.Post("/:id/logs", logHandler.CreateLog)

	// // ✅ Nested: Reports under Project
	// projects.Post("/:id/reports/generate", reportHandler.GenerateReport)
	// projects.Get("/:id/reports", reportHandler.GetProjectReports)
}