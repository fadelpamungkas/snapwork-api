package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	AssessmentEntity struct {
		Id     primitive.ObjectID `json:"_id,omitempty"`
		Item_1 string             `json:"item_1,omitempty"`
		Item_2 string             `json:"item_2,omitempty"`
		Item_3 string             `json:"item_3,omitempty"`
		Item_4 string             `json:"item_4,omitempty"`
	}
	AssessmentRequest struct {
		Id     primitive.ObjectID `json:"_id,omitempty"`
		Item_1 string             `json:"item_1,omitempty"`
		Item_2 string             `json:"item_2,omitempty"`
		Item_3 string             `json:"item_3,omitempty"`
		Item_4 string             `json:"item_4,omitempty"`
	}
	AssessmentResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
