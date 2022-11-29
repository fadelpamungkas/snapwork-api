package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AssessmentController struct {
	f       *fiber.App
	usecase usecase.AssessmentUsecaseI
}

func NewAssessmentController(f *fiber.App, usecase usecase.AssessmentUsecaseI) *AssessmentController {
	return &AssessmentController{
		f:       f,
		usecase: usecase,
	}
}

func (nc *AssessmentController) InsertAssessment(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.AssessmentRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AssessmentResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.AssessmentResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := nc.usecase.InsertAssessmentUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (nc *AssessmentController) GetAllAssessment(c *fiber.Ctx) error {

	data, err := nc.usecase.GetAllAssessmentUC(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.AssessmentResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all assessment",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}
