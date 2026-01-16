package routes

import (
	"devtracker/internal/handler"
	"devtracker/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RouteHandlers struct {
    AIInsight *handler.AIInsightHandler
    Project   *handler.ProjectHandler
    Milestone *handler.MilestoneHandler
    Log       *handler.LogHandler
    Report    *handler.ReportHandler
    User      *handler.UserHandler
}

func SetupRoutes(app *fiber.App, handlers RouteHandlers) {
    // API version 1
    api := app.Group("/api/v1")

    // Public routes
    setupAuthRoutes(api, handlers.User)

    // Protected routes (require authentication)
    protected := api.Group("", middleware.JWTProtected())

    // Setup domain-specific routes
	log.Printf("Project handler: %#v", handlers.Project)

    setupAIInsightRoutes(protected, handlers.AIInsight)
    setupProjectRoutes(protected, handlers.Project, handlers.Milestone, handlers.Log, handlers.Report)
    setupMilestoneRoutes(protected, handlers.Milestone)
    setupLogRoutes(protected, handlers.Log)
    setupReportRoutes(protected, handlers.Report)
    setupUserRoutes(protected, handlers.User)
}