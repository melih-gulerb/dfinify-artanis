package routes

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/repositories/definitionRepository"
	"artanis/src/repositories/projectUserRepository"
	"artanis/src/services"
	"github.com/gofiber/fiber/v2"
)

func SetupDefinitionRoutes(app *fiber.App, definitionRepository *definitionRepository.DefinitionRepository,
	projectUserRepo *projectUserRepository.ProjectUserRepository, cfg *configs.Config, definitionChangeService *services.DefinitionChangeService) {

	definitionHandler := handlers.NewDefinitionHandler(definitionRepository, projectUserRepo, definitionChangeService, cfg)

	definitionGroup := app.Group("/definitions")

	definitionGroup.Post("/", definitionHandler.Register)
	definitionGroup.Get("/:id", definitionHandler.Paginate)
	definitionGroup.Put("/", definitionHandler.Update)
	definitionGroup.Delete("/:id", definitionHandler.Delete)
}
