package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func NewsRoute(app *fiber.App, u usecase.NewsUsecaseI) {
	c := controllers.NewNewsController(app, u)

	api := app.Group("/api")

	api.Get("/news", c.GetAllNews)            // Get all companies
	api.Get("/news/:newsId", c.GetOneNews)    // Get all companies
	api.Post("/news", c.InsertNews)           // Get all companies
	api.Put("/news", c.UpdateNews)            // Update an existing post
	api.Delete("/news/:newsId", c.DeleteNews) // Delete post
}
