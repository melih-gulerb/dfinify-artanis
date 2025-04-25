package main

import (
	"artanis/src/clients"
	"artanis/src/configs"
	"artanis/src/middlewares"
	"artanis/src/routes"
	"artanis/src/services"
)

func main() {
	cfg := configs.InitConfig()
	db := configs.InitDB(cfg.MSSQLConnectionString)

	repositories := configs.InitDbContext(db)

	app := configs.InitFiber()
	app.Use(middlewares.PanicRecoveryMiddleware())

	slack := clients.NewSlackClient(cfg.SlackToken)
	definitionChangeService := services.NewDefinitionChangeService(repositories.DefinitionChangeRepository, slack)

	routes.SetupProjectRoutes(app, repositories.ProjectRepository, cfg)
	routes.SetupCollectionRoutes(app, repositories.CollectionRepository, repositories.ProjectUserRepository, cfg)
	routes.SetupProjectUserRoutes(app, repositories.ProjectUserRepository, cfg)
	routes.SetupDefinitionRoutes(app, repositories.DefinitionRepository, repositories.ProjectUserRepository, cfg, definitionChangeService)
	routes.SetupDefinitionChangeRoutes(app, db, cfg)

	_ = app.Listen(":4001")
}
