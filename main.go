package main

import (
	"artanis/src/configs"
	"artanis/src/middlewares"
	"artanis/src/routes"
)

func main() {
	cfg := configs.InitConfig()
	db := configs.InitDB(cfg.MSSQLConnectionString)
	configs.InitFiber()

	app := configs.InitFiber()
	app.Use(middlewares.AuthorizationMiddleware(cfg.JWTSecret))

	routes.SetupProjectRoutes(app, db, cfg)
	routes.SetupCollectionRoutes(app, db, cfg)
	routes.SetupProjectUserRoutes(app, db)
	routes.SetupDefinitionRoutes(app, db, cfg)

	_ = app.Listen(":4000")
}
