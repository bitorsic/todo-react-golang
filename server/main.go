package main

import (
	"os"
	"task-inator3000/config"
	"task-inator3000/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	config.DBConnect()
	config.RedisConnect()
}

func main() {
	app := fiber.New()

	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
