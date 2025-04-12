package routes

import (
	"database/sql"
	"divine-shield/src/handlers"
	"divine-shield/src/repositories/projectUserRepository"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectUserRoutes(app *fiber.App, db *sql.DB) {
	projectUserRepo := projectUserRepository.NewProjectUserRepository(db)

	projectUserHandler := handlers.NewProjectUserHandler(*projectUserRepo)

	projectUserGroup := app.Group("/project-users")

	projectUserGroup.Get("/assign", projectUserHandler.AssignUser)
}
