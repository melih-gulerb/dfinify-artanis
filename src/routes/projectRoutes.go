package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/projectRepository"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoutes(app *fiber.App, db *projectRepository.ProjectRepository, cfg *configs.Config) {
	projectHandler := handlers.NewProjectHandler(db, cfg)

	projectGroup := app.Group("/projects")

	projectGroup.Post("/", projectHandler.Register)
	projectGroup.Get("/", projectHandler.Paginate)
	projectGroup.Put("/", projectHandler.Update)
	projectGroup.Delete("/:id", projectHandler.Delete)
}
