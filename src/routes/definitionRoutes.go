package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/definitionRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupDefinitionRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	definitionRepo := definitionRepository.NewDefinitionRepository(db)

	definitionHandler := handlers.NewDefinitionHandler(definitionRepo, cfg)

	definitionGroup := app.Group("/definitions")

	definitionGroup.Post("/", definitionHandler.Register)
	definitionGroup.Get("/:id", definitionHandler.Paginate)
	definitionGroup.Put("/", definitionHandler.Update)
	definitionGroup.Delete("/:id", definitionHandler.Delete)
}
