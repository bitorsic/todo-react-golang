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
	//--------------------- NOTE TO SELF/CONTRIBUTOR ---------------------//
	// for the sake of consistency, when creating a controller,
	// if request fails => "error" (error should be in lowercase)
	// if request succeeds => directly return data OR just send status code
	// and use appropriate http status codes
	//--------------------------------------------------------------------//

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	// Use /api for all the api calls
	api := app.Group("/api")

	// auth
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/refresh", controllers.TokenRefresh)

	// routes below this need to be protected so using the auth middleware now
	api.Use(middleware.VerifyAuthToken)

	// tasklists and tasks
	tasks := api.Group("/tasks")

	// tasklists
	tasks.Get("/", controllers.GetTaskLists)
	tasks.Post("/", controllers.AddTaskList)

	// tasks
	tasks.Post("/:taskListID", controllers.AddTask)
}
