package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func PersonRoute(app *fiber.App, u usecase.PersonUsecaseI) {
	c := controllers.NewPersonController(app, u)

	api := app.Group("/api")

	api.Get("/person/:personId", c.GetPerson)                         // Get person
	api.Post("/person", c.InsertPerson)                               // create new person data
	api.Put("/person/notification", c.InsertNotification)             // Add new notification
	api.Put("/person", c.UpdatePerson)                                // Update an existing person
	api.Put("/person/document", c.UpdateDocumentPerson)               // Update Document
	api.Put("/person/portfolio", c.UpdatePortfolioPerson)             // Update Portfolio
	api.Put("/person/selfdevelopment", c.UpdateSelfDevelopmentPerson) // Update SelfDev
}
