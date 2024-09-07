package routes

import (
	"golang-backend/controllers"
	"golang-backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())

	// auth
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// homepage
	app.Get("/", middleware.VerifyJWT, controllers.Home)
}
