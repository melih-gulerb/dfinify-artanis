package routes

import (
	"artanis/src/clients"
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories/collectionRepository"
	"artanis/src/repositories/projectUserRepository"
	"github.com/gofiber/fiber/v2"
)

func SetupCollectionRoutes(app *fiber.App, collectionRepository *collectionRepository.CollectionRepository,
	projectUserRepository *projectUserRepository.ProjectUserRepository, cfg *configs.Config) {

	collectionHandler := handlers.NewCollectionHandler(collectionRepository, *projectUserRepository, cfg)

	collectionGroup := app.Group("/collections")

	divineShield := clients.NewDivineShieldClient(cfg.DivineShieldBaseUrl)
	collectionGroup.Use(middlewares.AuthorizationMiddleware(divineShield))

	collectionGroup.Post("/", collectionHandler.Register)
	collectionGroup.Get("/:id", collectionHandler.Paginate)
	collectionGroup.Put("/", collectionHandler.Update)
	collectionGroup.Delete("/:id", collectionHandler.Delete)
}
