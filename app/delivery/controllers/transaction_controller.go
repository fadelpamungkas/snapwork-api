package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	f           *fiber.App
	usecase usecase.TransactionUsecaseI
}

func NewTransactionController(f *fiber.App, usecase usecase.TransactionUsecaseI) *TransactionController {
	return &TransactionController{
		f:       f,
		usecase: usecase,
	}
}

func (uc *TransactionController) InsertOrder(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.OrderRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.OrderResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.OrderResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.InsertOrderUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *TransactionController) GetAllOrder(c *fiber.Ctx) error {

	data, err := uc.usecase.GetAllOrderUC(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.OrderResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all order data",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// authUser := c.Locals("authUser")
	return c.Status(data.Status).JSON(data)
}
