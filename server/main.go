package main

import (
	"golang-backend/config"
	"golang-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.DBConnect()
	routes.Setup(app)

	app.Listen(":3000")
}
