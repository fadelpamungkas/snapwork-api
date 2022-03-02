package controllers

import (
	"context"
	"golangapi/models/entities"
	"golangapi/models/requests"
	"golangapi/models/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(requests.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	validate := validator.New()
	if validationErr := validate.Struct(loginRequest); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}

	var user entities.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error getting user",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"token": "secret",
	})
}
