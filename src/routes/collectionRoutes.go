package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/collectionRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupCollectionRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	collectionRepo := collectionRepository.NewCollectionRepository(db)

	collectionHandler := handlers.NewCollectionHandler(collectionRepo, cfg)

	collectionGroup := app.Group("/collections")

	collectionGroup.Post("/collection", collectionHandler.Register)
	collectionGroup.Get("/collection", collectionHandler.Paginate)
	collectionGroup.Put("/collection", collectionHandler.Update)
	collectionGroup.Delete("/collection", collectionHandler.Delete)
}
