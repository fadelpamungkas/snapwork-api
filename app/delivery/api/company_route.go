package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func CompanyRoute(app *fiber.App, u usecase.CompanyUsecaseI) {
	c := controllers.NewCompanyController(app, u)

	api := app.Group("/api")

	// api.Get("/user/:userId", c.GetOneUser)                // Get a single user
	api.Get("/companies", c.GetAllCompanies)                   // Get all companies
	api.Get("/company/:companyId", c.GetCompany) // Get all companies
	api.Post("/company", c.InsertCompany)                      // Get all companies
	api.Post("/company/job", c.InsertJob)                      // Get all companies
}
