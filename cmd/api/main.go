package main

import (
	"devtracker/internal/handler"
	"devtracker/internal/repository/postgres"
	"devtracker/internal/routes"
	"devtracker/internal/service"
	"devtracker/pkg/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    // Test database connection
    log.Println("Starting DevTrackr API...")
    
    database := db.NewConnection()
    log.Println("Database connected successfully")

    // Initialize Fiber app
    app := fiber.New(fiber.Config{
        AppName: "DevTrackr API v1.0",
    })

    // initialize repo
    userRepository := postgres.NewUserPG(database);
    projectRepository := postgres.NewProjectPG(database);
    reportRepository := postgres.NewReportPG(database);
    logRepository := postgres.NewLogPG(database);
    aIInsightRepository := postgres.NewAIInsightPG(database);
    milestoneRepository := postgres.NewMilestonePG(database);






    // initialize service
    userService := service.NewUserService(userRepository)
    projectService := service.NewProjectService(projectRepository)
    reportService := service.NewReportService(reportRepository)
    logService := service.NewLogService(logRepository)
    aiInsightService := service.NewAIInsightService(aIInsightRepository)
    milestoneService := service.NewMilestoneService(milestoneRepository)





    // handler
    userHandler := handler.NewUserHandler(userService)
    projectHandler := handler.NewProjectHandler(projectService)
    reportHandler := handler.NewReportHandler(reportService)
    logHandler := handler.NewLogHandler(logService)
    aiInsightHandler := handler.NewAIInsightHandler(aiInsightService)
    milestoneHandler := handler.NewMilestoneHandler(milestoneService)




    handlers := routes.RouteHandlers{
		User:      userHandler,
        Project: projectHandler,
        Report: reportHandler,
        Log: logHandler,
        AIInsight: aiInsightHandler,
        Milestone: milestoneHandler,
	}

	routes.SetupRoutes(app, handlers)


	log.Println("API routes registered successfully")


    // Add basic middleware
   
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":    "up",
            "message":   "DevTrackr API is running",
            "database":  "connected",
            "timestamp": fiber.Map{},
        })
    })

    // Test database endpoint
    app.Get("/test-db", func(c *fiber.Ctx) error {
        sqlDB, err := database.DB()
        if err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error": "Failed to get database instance",
            })
        }

        if err := sqlDB.Ping(); err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error": "Database ping failed",
            })
        }

        return c.JSON(fiber.Map{
            "message": "Database connection is healthy",
            "status":  "ok",
        })
    })

    // Start server
    port := ":8080"
    log.Printf("🚀 Server starting on http://localhost%s", port)
    log.Printf("🏥 Health check: http://localhost%s/health", port)
    log.Printf("🗄️ Database test: http://localhost%s/test-db", port)
    
    log.Fatal(app.Listen(port))
}