package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserEntity struct {
		Id       primitive.ObjectID `json:"_id,omitempty"`
		Name     string             `json:"name,omitempty" validate:"required"`
		Email    string             `json:"email,omitempty" validate:"required,email"`
		Password string             `json:"password,omitempty" validate:"required"`
		Role     string             `json:"role,omitempty""`
	}
	UserRequest struct {
		Id       primitive.ObjectID `json:"_id,omitempty"`
		Name     string             `json:"name" validate:"required"`
		Email    string             `json:"email" validate:"required,email"`
		Password string             `json:"password" validate:"required"`
		Role     string             `json:"role,omitempty""`
	}
	UserResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
