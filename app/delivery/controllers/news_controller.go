package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type NewsController struct {
	f       *fiber.App
	usecase usecase.NewsUsecaseI
}

func NewNewsController(f *fiber.App, usecase usecase.NewsUsecaseI) *NewsController {
	return &NewsController{
		f:       f,
		usecase: usecase,
	}
}

func (nc *NewsController) InsertNews(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.NewsRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewsResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewsResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := nc.usecase.InsertNewsUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (nc *NewsController) GetAllNews(c *fiber.Ctx) error {

	data, err := nc.usecase.GetAllNewsUC(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewsResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all news",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// authUser := c.Locals("authUser")
	return c.Status(data.Status).JSON(data)
}

func (pc *NewsController) GetOneNews(c *fiber.Ctx) error {
	newsId := c.Params("newsId")
	data, err := pc.usecase.GetOneNewsUC(context.Background(), newsId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewsResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting news",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}

func (nc *NewsController) UpdateNews(c *fiber.Ctx) error {
	var req models.NewsRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewsResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	data, err := nc.usecase.UpdateNewsUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (nc *NewsController) DeleteNews(c *fiber.Ctx) error {
	newsId := c.Params("newsId")
	data, err := nc.usecase.DeleteNewsUC(context.Background(), newsId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.NewsResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error delete news",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data).JSON(data)
}
