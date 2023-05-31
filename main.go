package main

import (
	"rest-go/database"
	"rest-go/database/migration"
	"rest-go/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()

	// INITIAL MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	// EXPORT PORT
	app.Listen(":3000")
}
