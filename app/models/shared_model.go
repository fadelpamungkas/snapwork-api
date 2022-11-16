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
	Education struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		Name          string             `json:"name,omitempty"`
		Concentration string             `json:"concentration,omitempty"`
		Date          string             `json:"date,omitempty"`
		Degree        string             `json:"degree,omitempty"`
	}
	SelfDevelopment struct {
		Id     primitive.ObjectID `json:"_id,omitempty"`
		Score  string             `json:"score,omitempty"`
		Status string             `json:"status,omitempty"`
		File   string             `json:"file,omitempty"`
	}
	Career struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		ApplicationId primitive.ObjectID `json:"applicationId,omitempty"`
		Status        string             `json:"name,omitempty"`
		CompanyName   string             `json:"concentration,omitempty"`
		Date          string             `json:"date,omitempty"`
		Position      string             `json:"position,omitempty"`
		Type          string             `json:"type,omitempty"`
	}
)
