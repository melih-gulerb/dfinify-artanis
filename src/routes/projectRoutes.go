package routes

import (
	"artanis/src/clients"
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories/projectRepository"
	"github.com/gofiber/fiber/v2"
)

func SetupProjectRoutes(app *fiber.App, db *projectRepository.ProjectRepository, cfg *configs.Config) {
	projectHandler := handlers.NewProjectHandler(db, cfg)

	projectGroup := app.Group("/projects")

	divineShield := clients.NewDivineShieldClient(cfg.DivineShieldBaseUrl)
	projectGroup.Use(middlewares.AuthorizationMiddleware(divineShield))

	projectGroup.Post("/", projectHandler.Register)
	projectGroup.Get("/", projectHandler.Paginate)
	projectGroup.Put("/", projectHandler.Update)
	projectGroup.Delete("/:id", projectHandler.Delete)
	projectGroup.Post("/secret/:id", projectHandler.GenerateSecret)

	projectFeedGroup := app.Group("/projects-feed")

	projectFeedGroup.Get("/:id", projectHandler.GetProjectFeed)
}
