package main

import (
	"artanis/src/configs"
	"artanis/src/handlers"
	"artanis/src/middlewares"
	"artanis/src/repositories"
)

func main() {
	cfg := configs.InitConfig()
	db := configs.InitDB(cfg.MSSQLConnectionString)
	configs.InitFiber()

	projectRepository := repositories.NewProjectRepository(db)
	projectHandler := handlers.NewProjectHandler(projectRepository, cfg)
	collectionRepository := repositories.NewCollectionRepository(db)
	collectionHandler := handlers.NewCollectionHandler(collectionRepository, cfg)
	definitionRepository := repositories.NewDefinitionRepository(db)
	definitionHandler := handlers.NewDefinitionHandler(definitionRepository, cfg)

	app := configs.InitFiber()
	app.Use(middlewares.AuthorizationMiddleware(cfg.JWTSecret))
	app.Post("/project", projectHandler.Register)
	app.Get("/project", projectHandler.Paginate)
	app.Put("/project", projectHandler.Update)
	app.Delete("/project", projectHandler.Delete)

	app.Post("/collection", collectionHandler.Register)
	app.Get("/collection", collectionHandler.Paginate)
	app.Put("/collection", collectionHandler.Update)
	app.Delete("/collection", collectionHandler.Delete)

	app.Post("/definition", definitionHandler.Register)
	app.Get("/definition", definitionHandler.Paginate)
	app.Put("/definition", definitionHandler.Update)
	app.Delete("/definition", definitionHandler.Delete)

	_ = app.Listen(":4000")
}
