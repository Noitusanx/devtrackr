package routes

import (
	"devtracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func setupReportRoutes(app fiber.Router, handler *handler.ReportHandler){

	report := app.Group("/reports");

	report.Post("/", handler.CreateReport);
	report.Get("/:id", handler.GetLogByID);
	report.Put("/:id", handler.UpdateReport);
	report.Delete("/:id", handler.DeleteReport);


	report.Get("/projects/:projectID", handler.GetReportsByProject)
}