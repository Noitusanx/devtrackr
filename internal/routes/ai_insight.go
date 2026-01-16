package routes

import (
	"devtracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func setupAIInsightRoutes(app fiber.Router, handler *handler.AIInsightHandler) {

	insights := app.Group("/insights")

	insights.Post("/", handler.CreateInsight)
	insights.Get("/project/:projectID", handler.GetInsightsByProject)
	insights.Get("/:id", handler.GetInsightByID)


	// insights.Put("/:id", handler.UpdateInsight)
	// insights.Delete("/:id", handler.DeleteInsight)
	// insights.Get("/project/:projectID/type/:insightType", handler.GetInsightByProjectAndTypeAndDate)
}