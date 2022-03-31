package entities

import (
	// "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id         primitive.ObjectID `json:"_id,omitempty"`
	Title      string             `json:"title,omitempty" validate:"required"`
	Content    string             `json:"content,omitempty" validate:"required"`
	Category   string             `json:"category,omitempty" validate:"required"`
	AuthorId   primitive.ObjectID `json:"authorId,omitempty" validate:"required"`
	AuthorName string             `json:"authorName,omitempty" validate:"required"`
	Images     []Image            `json:"image,omitempty"`
	// Tier       struct {
	// 	Silver   TierContent `json:"silver,omitempty" validate:"required"`
	// 	Gold     TierContent `json:"gold,omitempty" validate:"required"`
	// 	Platinum TierContent `json:"platinum,omitempty" validate:"required"`
	// } `json:"tier,omitempty" validate:"required"`
}

type TierContent struct {
	Description string   `json:"description,omitempty" validate:"required"`
	Price       float64  `json:"price,omitempty" validate:"required"`
	Offer       []string `json:"offer,omitempty" validate:"required"`
}

type Image struct {
	Id   primitive.ObjectID `json:"_id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required"`
	Url  string             `json:"url,omitempty"`
}
