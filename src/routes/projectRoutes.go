package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/projectRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	projectRepo := projectRepository.NewProjectRepository(db)

	projectHandler := handlers.NewProjectHandler(projectRepo, cfg)

	projectGroup := app.Group("/projects")

	projectGroup.Post("/", projectHandler.Register)
	projectGroup.Get("/", projectHandler.Paginate)
	projectGroup.Put("/", projectHandler.Update)
	projectGroup.Delete("/:id", projectHandler.Delete)
}
