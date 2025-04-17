package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/collectionRepository"
	"artanis/src/repositories/projectUserRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupCollectionRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	collectionRepo := collectionRepository.NewCollectionRepository(db)
	projectUserRepo := projectUserRepository.NewProjectUserRepository(db)

	collectionHandler := handlers.NewCollectionHandler(collectionRepo, *projectUserRepo, cfg)

	collectionGroup := app.Group("/collections")

	collectionGroup.Post("/", collectionHandler.Register)
	collectionGroup.Get("/:id", collectionHandler.Paginate)
	collectionGroup.Put("/", collectionHandler.Update)
	collectionGroup.Delete("/:id", collectionHandler.Delete)
}
