package routes

import (
	"golangapi/configs"
	"golangapi/controllers"
	"golangapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Static("/public", configs.BasePath+"/public/assets", fiber.Static{
		Browse:   true,
		Compress: false,
	})

	app.Get("/api/users", middleware.Auth, controllers.GetAllUsers) // Get all users
	app.Get("/api/user/:userId", controllers.GetSingleUser)         // Get a single user
	app.Post("/api/users", controllers.CreateUser)                  // Create a new user
	app.Put("/api/user/:userId", controllers.UpdateUser)            // Update an existing user
	app.Delete("/api/user/:userId", controllers.DeleteUser)         // Delete user
	app.Post("/login", controllers.Login)                           // login authentication
}
