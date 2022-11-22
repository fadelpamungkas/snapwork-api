package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(app *fiber.App, u usecase.TransactionUsecaseI) {
	c := controllers.NewTransactionController(app, u)

	api := app.Group("/api")

	// api.Get("/users", middlewares.Auth, c.GetAllUsers)    // Get all users
	// api.Get("/authuser", middlewares.Auth, c.GetAuthUser) // Get auth user
	api.Get("/transaction/orders", c.GetAllOrder)                         // Get all order
	api.Post("/transaction/order", c.InsertOrder)                         // Create new order
	api.Post("/transaction/application", c.InsertApplication)             // Insert new application
	api.Put("/transaction/application/status", c.UpdateApplicationStatus) // Update application status
}
