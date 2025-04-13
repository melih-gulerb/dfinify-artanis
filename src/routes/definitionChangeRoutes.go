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

	definitionChangeGroup.Post("/definitionChange", definitionChangeHandler.Register)
	definitionChangeGroup.Get("/definitionChange", definitionChangeHandler.Paginate)
	definitionChangeGroup.Put("/definitionChange", definitionChangeHandler.Update)
}
