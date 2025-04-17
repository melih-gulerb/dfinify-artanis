package configs

import (
	"artanis/src/logging"
	"artanis/src/middlewares"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gofiber/fiber/v2"
)

func InitDB(dsn string) *sql.DB {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		logging.Log(logging.PANIC, fmt.Sprintf("Failed to connect DB %v", err))
	}

	if err = db.Ping(); err != nil {
		logging.Log(logging.PANIC, fmt.Sprintf("Failed to ping DB: %v", err))
	}

	logging.Log(logging.INFO, "MSSQL connection successfully established")

	return db
}

func InitFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.CustomErrorHandler,
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Artanis is running")
	})

	return app
}
