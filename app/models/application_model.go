package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ApplicationEntity struct {
		Id           primitive.ObjectID `json:"_id,omitempty"`
		UserId       primitive.ObjectID `json:"userid,omitempty"`
		CompanyId    primitive.ObjectID `json:"companyid,omitempty"`
		CompanyJobId primitive.ObjectID `json:"companyjobid,omitempty"`
		Status       string             `json:"status,omitempty"`
		CreatedAt    string             `json:"created_at,omitempty"`
	}
	ApplicationRequest struct {
		Id           primitive.ObjectID `json:"_id,omitempty"`
		UserId       primitive.ObjectID `json:"userid,omitempty"`
		CompanyId    primitive.ObjectID `json:"companyid,omitempty"`
		CompanyJobId primitive.ObjectID `json:"companyjobid,omitempty"`
		Status       string             `json:"status,omitempty"`
		CreatedAt    string             `json:"created_at,omitempty"`
	}
	ApplicationResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
