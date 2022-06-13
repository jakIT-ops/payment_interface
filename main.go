package main

import (
	"interface/database"
	"interface/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// INITIAL DATABASE
	database.ConnectDb()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":3000")

}
