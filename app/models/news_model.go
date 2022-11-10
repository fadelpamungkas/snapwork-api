package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	NewsEntity struct {
		Id        primitive.ObjectID `json:"_id,omitempty"`
		Author    string             `json:"author,omitempty"`
		Type      string             `json:"type,omitempty"`
		Title     string             `json:"title,omitempty"`
		Body      string             `json:"body,omitempty"`
		Quotes    string             `json:"quotes,omitempty"`
		Header    string             `json:"header,omitempty"`
		CreatedAt string             `json:"created_at,omitempty"`
	}
	NewsRequest struct {
		Id        primitive.ObjectID `json:"_id,omitempty"`
		Author    string             `json:"author,omitempty"`
		Type      string             `json:"type,omitempty"`
		Title     string             `json:"title,omitempty"`
		Body      string             `json:"body,omitempty"`
		Quotes    string             `json:"quotes,omitempty"`
		Header    string             `json:"header,omitempty"`
		CreatedAt string             `json:"created_at,omitempty"`
	}
	NewsResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
