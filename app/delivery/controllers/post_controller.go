package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	f       *fiber.App
	usecase usecase.PostUsecaseI
}

func NewPostController(f *fiber.App, usecase usecase.PostUsecaseI) *PostController {
	return &PostController{
		f:       f,
		usecase: usecase,
	}
}

func (pc *PostController) GetAllPost(c *fiber.Ctx) error {
	query := new(models.Query)

	// Queries are optional
	c.QueryParser(query)
	log.Println(*query)

	data, err := pc.usecase.GetAllUC(context.Background(), *query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all users",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	return c.Status(data.Status).JSON(data)
}

func (pc *PostController) GetAllPostByUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	data, err := pc.usecase.GetAllByUserUC(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all users",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	return c.Status(data.Status).JSON(data)
}

func (pc *PostController) GetOnePost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	data, err := pc.usecase.GetOneUC(context.Background(), postId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting post",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}

func (pc *PostController) InsertPost(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.PostRequest

	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PostResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PostResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PostResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid Multipart form",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	req.Images = form

	data, err := pc.usecase.InsertUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (pc *PostController) UpdatePost(c *fiber.Ctx) error {
	var req models.PostRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.PostResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	data, err := pc.usecase.UpdateUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (pc *PostController) DeletePost(c *fiber.Ctx) error {
	postId := c.Params("postId")
	data, err := pc.usecase.DeleteUC(context.Background(), postId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error delete post",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data).JSON(data)
}
