package routes

import (
	"artanis/src/handlers"
	"artanis/src/repositories/projectUserRepository"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectUserRoutes(app *fiber.App, db *projectUserRepository.ProjectUserRepository) {
	projectUserHandler := handlers.NewProjectUserHandler(*db)

	projectUserGroup := app.Group("/project-users")

	projectUserGroup.Get("/assign", projectUserHandler.AssignUser)
}
