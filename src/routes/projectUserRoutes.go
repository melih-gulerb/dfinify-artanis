package routes

import (
	"artanis/src/clients"
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories/projectUserRepository"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectUserRoutes(app *fiber.App, db *projectUserRepository.ProjectUserRepository, cfg *configs.Config) {
	projectUserHandler := handlers.NewProjectUserHandler(*db)

	projectUserGroup := app.Group("/project-users")

	divineShield := clients.NewDivineShieldClient(cfg.DivineShieldBaseUrl)
	projectUserGroup.Use(middlewares.AuthorizationMiddleware(divineShield))

	projectUserGroup.Post("/", projectUserHandler.Register)
	projectUserGroup.Put("/:id", projectUserHandler.UpdateProjectUserRole)
	projectUserGroup.Delete("/:id", projectUserHandler.Delete)
	projectUserGroup.Delete("/:id", projectUserHandler.Paginate)
}
