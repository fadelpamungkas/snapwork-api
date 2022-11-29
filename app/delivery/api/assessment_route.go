package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func AssessmentRoute(app *fiber.App, u usecase.AssessmentUsecaseI) {
	c := controllers.NewAssessmentController(app, u)

	api := app.Group("/api")

	api.Get("/assessment", c.GetAllAssessment) // Get all companies
	// api.Get("/news/:newsId", c.GetOneAssessment)    // Get all companies
	api.Post("/assessment", c.InsertAssessment) // Get all companies
	// api.Put("/news", c.UpdateAssessment)            // Update an existing post
	// api.Delete("/news/:newsId", c.DeleteAssessment) // Delete post
}
