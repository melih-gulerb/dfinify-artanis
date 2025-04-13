package main

import (
	"artanis/src/clients"
	"artanis/src/configs"
	"artanis/src/middlewares"
	"artanis/src/routes"
)

func main() {
	cfg := configs.InitConfig()
	db := configs.InitDB(cfg.MSSQLConnectionString)

	divineShield := clients.NewDivineShieldClient(cfg.DivineShieldBaseUrl)

	app := configs.InitFiber()
	app.Use(middlewares.AuthorizationMiddleware(divineShield))

	routes.SetupProjectRoutes(app, db, cfg)
	routes.SetupCollectionRoutes(app, db, cfg)
	routes.SetupProjectUserRoutes(app, db)
	routes.SetupDefinitionRoutes(app, db, cfg)
	routes.SetupDefinitionChangeRoutes(app, db, cfg)

	_ = app.Listen(":4001")
}
