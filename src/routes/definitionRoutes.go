package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories/definitionRepository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupDefinitionRoutes(app *fiber.App, db *sql.DB, cfg *configs.Config) {
	definitionRepo := definitionRepository.NewDefinitionRepository(db)

	definitionHandler := handlers.NewDefinitionHandler(definitionRepo, cfg)

	definitionGroup := app.Group("/definitions")

	definitionGroup.Use(middlewares.AuthorizationMiddleware(cfg.JWTSecret))
	definitionGroup.Post("/definition", definitionHandler.Register)
	definitionGroup.Get("/definition", definitionHandler.Paginate)
	definitionGroup.Put("/definition", definitionHandler.Update)
	definitionGroup.Delete("/definition", definitionHandler.Delete)
}
