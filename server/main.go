package main

import (
	"golang-backend/config"
	"golang-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.DBConnect()
	config.RedisConnect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
