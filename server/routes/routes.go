package routes

import (
	"golang-backend/controllers"
	"golang-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", middleware.VerifyJWT, controllers.Home)
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
