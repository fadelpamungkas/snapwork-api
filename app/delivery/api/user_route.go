package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/delivery/controllers/middlewares"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, u usecase.UserUsecaseI) {
	c := controllers.NewUserController(app, u)

	app.Get("/api/users", middlewares.Auth, c.GetAllUsers)    // Get all users
	app.Get("/api/authuser", middlewares.Auth, c.GetAuthUser) // Get auth user
	app.Get("/api/user/:userId", c.GetOneUser)                // Get a single user
	app.Post("/api/user", c.InsertUser)                       // Create a new user
	app.Put("/api/user", c.UpdateUser)                        // Update an existing user
	app.Delete("/api/user/:userId", c.DeleteUser)             // Delete user
	app.Post("/login", c.Login)                               // login authentication
}
