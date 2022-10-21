package controllers

import (
	"context"
	"golangapi/app/models"
	"golangapi/app/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	f       *fiber.App
	usecase usecase.UserUsecaseI
}

func NewUserController(f *fiber.App, usecase usecase.UserUsecaseI) *UserController {
	return &UserController{
		f:       f,
		usecase: usecase,
	}
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	var user models.LoginRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.UserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	validate := validator.New()
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.UserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}

	data, err := uc.usecase.LoginUC(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.UserResponse{
			Status:  fiber.StatusNotFound,
			Message: "Account Not Found",
			Data: &fiber.Map{
				"data": err,
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {

	data, err := uc.usecase.GetAllUC(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all users",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	authUser := c.Locals("authUser")
	return c.Status(data.Status).JSON(authUser)
}

func (uc *UserController) GetAuthUser(c *fiber.Ctx) error {
	authUser := c.Locals("authUser")
	if authUser == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting users",
			Data: &fiber.Map{
				"data": "",
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		models.UserResponse{
			Status:  fiber.StatusOK,
			Message: "Successfully get user",
			Data: &fiber.Map{
				"data": authUser,
			},
		})
}

func (uc *UserController) GetOneUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	data, err := uc.usecase.GetOneUserUC(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting all users",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data.Status).JSON(data)
}

func (uc *UserController) InsertUser(c *fiber.Ctx) error {
	var validate = validator.New()
	var req models.UserRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.UserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	//use validator library to validate required fields
	if validationErr := validate.Struct(&req); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.UserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}
	data, err := uc.usecase.InsertUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	var req models.UserRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.UserResponse{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}
	data, err := uc.usecase.UpdateUC(context.Background(), req)
	if err != nil {
		return err
	}

	return c.Status(data).JSON(data)
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	data, err := uc.usecase.DeleteUC(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error delete user",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(data).JSON(data)
}
