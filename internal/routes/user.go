package routes

import (
	"devtracker/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(app fiber.Router, handler *handler.UserHandler){
	auth := app.Group("/auth");

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}

func setupUserRoutes(app fiber.Router, handler *handler.UserHandler)  {
	user := app.Group("/users")

	user.Get("/me", handler.GetCurrentUser)
	user.Put("/me", handler.UpdateProfile)
	// user.Get("/dashboard", handler.GetDashboard)
	user.Delete("/me", handler.DeleteAccount)
}