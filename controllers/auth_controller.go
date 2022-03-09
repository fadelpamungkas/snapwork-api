package controllers

import (
	"context"
	"golangapi/models/entities"
	"golangapi/models/requests"
	"golangapi/models/responses"
	"golangapi/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(requests.LoginRequest)

	var user entities.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	err := userCollection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&user)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "Wrong Credential",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// validating password
	errPass := utils.CheckHashPassword(loginRequest.Password, user.Password)

	if errPass != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "Wrong Credential",
			Data: &fiber.Map{
				"data": errPass.Error(),
			},
		})
	}

	// generate jwt token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id.Hex()
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerate := utils.GenerateToken(&claims)

	if errGenerate != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error generating token",
			Data: &fiber.Map{
				"data": errGenerate.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
