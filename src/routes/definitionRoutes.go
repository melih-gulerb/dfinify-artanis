package routes

import (
	"artanis/src/clients"
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories/definitionRepository"
	"artanis/src/repositories/projectUserRepository"
	"artanis/src/services"
	"github.com/gofiber/fiber/v2"
)

func SetupDefinitionRoutes(app *fiber.App, definitionRepository *definitionRepository.DefinitionRepository,
	projectUserRepo *projectUserRepository.ProjectUserRepository, cfg *configs.Config, definitionChangeService *services.DefinitionChangeService) {

	definitionHandler := handlers.NewDefinitionHandler(definitionRepository, projectUserRepo, definitionChangeService, cfg)

	definitionGroup := app.Group("/definitions")

	divineShield := clients.NewDivineShieldClient(cfg.DivineShieldBaseUrl)
	definitionGroup.Use(middlewares.AuthorizationMiddleware(divineShield))

	definitionGroup.Post("/", definitionHandler.Register)
	definitionGroup.Get("/:id", definitionHandler.Paginate)
	definitionGroup.Put("/name", definitionHandler.UpdateName)
	definitionGroup.Put("/value", definitionHandler.UpdateValue)
	definitionGroup.Put("/state", definitionHandler.UpdateState)
	definitionGroup.Delete("/:id", definitionHandler.Delete)
}
