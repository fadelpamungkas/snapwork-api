package models

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PostEntity struct {
		Id         primitive.ObjectID `json:"_id,omitempty"`
		Title      string             `json:"title,omitempty" validate:"required"`
		Content    string             `json:"content,omitempty" validate:"required"`
		Category   string             `json:"category,omitempty" validate:"required"`
		Price      float64            `json:"price,omitempty" validate:"required"`
		AuthorId   primitive.ObjectID `json:"authorId,omitempty" validate:"required"`
		AuthorName string             `json:"authorName,omitempty" validate:"required"`
		Images     []Image            `json:"images,omitempty"`
	}
	PostRequest struct {
		Id         primitive.ObjectID `json:"_id,omitempty"`
		Title      string             `json:"title" validate:"required"`
		Content    string             `json:"content" validate:"required"`
		Category   string             `json:"category" validate:"required"`
		Price      float64            `json:"price,omitempty" validate:"required"`
		AuthorId   primitive.ObjectID `json:"authorId" validate:"required"`
		AuthorName string             `json:"authorName" validate:"required"`
		Images     *multipart.Form    `json:"images,omitempty"`
	}
	PostResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
