package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	f       *fiber.App
	usecase usecase.CompanyUsecaseI
}

func NewCompanyController(f *fiber.App, usecase usecase.CompanyUsecaseI) *CompanyController {
	return &CompanyController{
		f:       f,
		usecase: usecase,
	}
}

func (uc *CompanyController) InsertCompany(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.CompanyRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CompanyResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CompanyResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.InsertCompanyUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *CompanyController) GetAllCompanies(c *fiber.Ctx) error {

	data, err := uc.usecase.GetAllCompaniesUC(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.CompanyResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all companies",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// authUser := c.Locals("authUser")
	return c.Status(data.Status).JSON(data)
}

func (uc *CompanyController) InsertJob(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.CompanyJobRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CompanyResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.CompanyJobResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.InsertJobUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *CompanyController) GetCompany(c *fiber.Ctx) error {
	companyId := c.Params("companyId")
	data, err := uc.usecase.GetCompany(context.Background(), companyId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.CompanyResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting company",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}

func (uc *CompanyController) GetJobCompany(c *fiber.Ctx) error {
	companyId := c.Params("companyId")
	jobId := c.Params("jobId")
	data, err := uc.usecase.GetJobCompany(context.Background(), companyId, jobId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting job company",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}
