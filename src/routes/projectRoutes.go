package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories/projectRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	projectRepo := projectRepository.NewProjectRepository(db)

	projectHandler := handlers.NewProjectHandler(projectRepo, cfg)

	projectGroup := app.Group("/projects")

	projectGroup.Use(middlewares.AuthorizationMiddleware(cfg.JWTSecret))
	projectGroup.Post("/project", projectHandler.Register)
	projectGroup.Get("/project", projectHandler.Paginate)
	projectGroup.Put("/project", projectHandler.Update)
	projectGroup.Delete("/project", projectHandler.Delete)
}
