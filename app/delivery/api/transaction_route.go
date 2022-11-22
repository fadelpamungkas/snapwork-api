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
	api.Get("/transaction/orders", c.GetAllOrder)                                              // Get all companies
	api.Post("/transaction/order", c.InsertOrder)                                              // Get all companies
	api.Get("/transaction/applicationsbycompanyid/:companyId", c.GetAllApplicationByCompanyId) // Get all companies
	api.Get("/transaction/applicationsbyuserid/:userId", c.GetAllApplicationByUserId)          // Get all companies
	api.Post("/transaction/application", c.InsertApplication)                                  // Get all companies
	api.Put("/transaction/application/status", c.UpdateApplicationStatus)                      // Update an existing user
}
