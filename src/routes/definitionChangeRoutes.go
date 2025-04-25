package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/definitionChangeRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupDefinitionChangeRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	definitionChangeRepo := definitionChangeRepository.NewDefinitionChangeRepository(db)

	definitionChangeHandler := handlers.NewDefinitionChangeHandler(definitionChangeRepo)

	definitionChangeGroup := app.Group("/definitionChanges")

	definitionChangeGroup.Get("/", definitionChangeHandler.Paginate)
	definitionChangeGroup.Put("/", definitionChangeHandler.Update)
}
