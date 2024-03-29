package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PersonController struct {
	f       *fiber.App
	usecase usecase.PersonUsecaseI
}

func NewPersonController(f *fiber.App, usecase usecase.PersonUsecaseI) *PersonController {
	return &PersonController{
		f:       f,
		usecase: usecase,
	}
}

func (nc *PersonController) InsertPerson(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.PersonRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := nc.usecase.InsertPersonUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *PersonController) GetAllPerson(c *fiber.Ctx) error {

	data, err := uc.usecase.GetAllPersonUC(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.PersonResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all person",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// authUser := c.Locals("authUser")
	return c.Status(data.Status).JSON(data)
}

func (pc *PersonController) GetPerson(c *fiber.Ctx) error {
	newsId := c.Params("personId")
	data, err := pc.usecase.GetPersonUC(context.Background(), newsId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.PersonResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting person",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}

func (nc *PersonController) UpdatePerson(c *fiber.Ctx) error {
	var req models.PersonRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	data, err := nc.usecase.UpdatePersonUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *PersonController) InsertNotification(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.Notification
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.InsertNotificationUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *PersonController) UpdateDocumentPerson(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.PersonDocumentRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.UpdateDocumentPersonUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *PersonController) UpdatePortfolioPerson(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.Portfolio
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.UpdatePortfolioPersonUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *PersonController) UpdateSelfDevelopmentPerson(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.SelfDevelopment
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.UpdateSelfDevelopmentPersonUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (nc *PersonController) UpdateSelfDevelopmentPaymentPerson(c *fiber.Ctx) error {
	var req models.SelfDevelopmentPaymentRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PersonResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	data, err := nc.usecase.UpdateSelfDevelopmentPaymentPersonUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}
