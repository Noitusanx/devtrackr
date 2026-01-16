package routes

import (
	"devtracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func setupLogRoutes(app fiber.Router, handler *handler.LogHandler) {

	logs := app.Group("/logs")

	logs.Post("/", handler.CreateLog)
	logs.Get("/:id", handler.GetLogByID)
	logs.Put("/:id", handler.UpdateLog)
	logs.Delete("/:id", handler.DeleteLog)

	logs.Get("/user/:userID", handler.GetLogsByUser)
	logs.Get("/project/:projectID", handler.GetLogsByProject)

}