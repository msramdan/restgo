package main

import (
	"rest-go/database"
	"rest-go/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	app := fiber.New()
	// inital route
	route.RouteInit(app)
	app.Listen(":3000")
}
