package routes

import (
	"artanis/src/handlers"
	"artanis/src/repositories/projectUserRepository"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupProjectUserRoutes(app *fiber.App, db *sql.DB) {
	projectUserRepo := projectUserRepository.NewProjectUserRepository(db)

	projectUserHandler := handlers.NewProjectUserHandler(*projectUserRepo)

	projectUserGroup := app.Group("/project-users")

	projectUserGroup.Get("/assign", projectUserHandler.AssignUser)
}
