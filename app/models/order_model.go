package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	OrderEntity struct {
		Id        primitive.ObjectID `json:"_id,omitempty"`
		UserId    primitive.ObjectID `json:"userid,omitempty"`
		Status    string             `json:"status,omitempty"`
		Name      string             `json:"name,omitempty"`
		Amount    string             `json:"amount,omitempty"`
		FileProof string             `json:"fileproof,omitempty"`
		CreatedAt string             `json:"created_at,omitempty"`
	}
	OrderRequest struct {
		UserId    primitive.ObjectID `json:"userid,omitempty"`
		Status    string             `json:"status,omitempty"`
		Name      string             `json:"name,omitempty"`
		Amount    string             `json:"amount,omitempty"`
		FileProof string             `json:"fileproof,omitempty"`
		CreatedAt string             `json:"created_at,omitempty"`
	}
	OrderResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
