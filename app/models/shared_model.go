package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Image struct {
		Id   primitive.ObjectID `json:"_id,omitempty"`
		Name string             `json:"name,omitempty"`
		Url  string             `json:"url,omitempty"`
	}
)
