package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/delivery/controllers/middlewares"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, u usecase.UserUsecaseI) {
	c := controllers.NewUserController(app, u)

	api := app.Group("/api")

	api.Get("/users", middlewares.Auth, c.GetAllUsers)    // Get all users
	api.Get("/authuser", middlewares.Auth, c.GetAuthUser) // Get auth user
	api.Get("/user/:userId", c.GetOneUser)                // Get a single user
	// api.Get("/companies", c.GetAllCompanies)              // Get all companies
	// api.Post("/company", c.InsertCompany)                 // Get all companies
	api.Post("/user", c.InsertUser)           // Create a new user
	api.Put("/user/role", c.UpdateRole)       // Update an existing user
	api.Delete("/user/:userId", c.DeleteUser) // Delete user
	api.Post("/login", c.Login)               // Login authentication
}
