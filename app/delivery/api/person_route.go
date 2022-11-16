package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func PersonRoute(app *fiber.App, u usecase.PersonUsecaseI) {
	c := controllers.NewPersonController(app, u)

	api := app.Group("/api")

	api.Get("/person/:personId", c.GetPerson) // Get person
	api.Post("/person", c.InsertPerson)       // create new person data
	api.Put("/person", c.UpdatePerson)        // Update an existing person
}
