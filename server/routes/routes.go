package routes

import (
	"golang-backend/controllers"
	"golang-backend/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	// Use /api for all the api calls
	api := app.Group("/api")

	// auth
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// routes below this need to be protected so using the auth middleware now
	api.Use(middleware.VerifyJWT)

	// tasks
	api.Get("/tasks", controllers.GetTaskLists)
	api.Post("/tasks", controllers.AddTaskList)
	api.Post("/tasks/:taskListID", controllers.AddTask)
}
