package entities

import (
	// "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	Title    string             `json:"title,omitempty" validate:"required"`
	Content  string             `json:"content,omitempty" validate:"required"`
	Category string             `json:"category,omitempty" validate:"required"`
	Price    float64            `json:"price,omitempty" validate:"required"`
	// Photo      *fiber.Map         `json:"photo,omitempty" validate:"required"`
	// AuthorId   primitive.ObjectID `json:"authorId,omitempty" validate:"required"`
	AuthorName string `json:"authorName,omitempty" validate:"required"`
}
