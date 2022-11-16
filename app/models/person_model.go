package models

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PersonEntity struct {
		Id              primitive.ObjectID `json:"_id,omitempty"`
		UserId          primitive.ObjectID `json:"userid,omitempty"`
		Name            string             `json:"name,omitempty"`
		Birth           string             `json:"birth,omitempty"`
		Gender          string             `json:"gender,omitempty"`
		Religion        string             `json:"religion,omitempty"`
		Marriage        string             `json:"marriage,omitempty"`
		Hobby           string             `json:"hobby,omitempty"`
		Telephone       string             `json:"telephone,omitempty"`
		Email           string             `json:"email,omitempty"`
		Twitter         string             `json:"twitter,omitempty"`
		Linkedin        string             `json:"linkedin,omitempty"`
		Address         string             `json:"address,omitempty"`
		City            string             `json:"city,omitempty"`
		Province        string             `json:"province,omitempty"`
		State           string             `json:"state,omitempty"`
		About           string             `json:"about,omitempty"`
		Portfolio       string             `json:"portfolio,omitempty"`
		Avatar          Image              `json:"avatar,omitempty"`
		KTP             Image              `json:"ktp,omitempty"`
		Ijazah          Image              `json:"ijazah,omitempty"`
		SKCK            Image              `json:"skck,omitempty"`
		CV              Image              `json:"cv,omitempty"`
		Certificate     Image              `json:"certificate,omitempty"`
		SelfDevelopment SelfDevelopment    `json:"selfdevelopment,omitempty"`
		Education       Education          `json:"lasteducation,omitempty"`
		Career          []Career           `json:"career,omitempty"`
		CreatedAt       string             `json:"created_at,omitempty"`
		UpdatedAt       string             `json:"updated_at,omitempty"`
	}
	PersonRequest struct {
		Id              primitive.ObjectID `json:"_id,omitempty"`
		UserId          primitive.ObjectID `json:"userid,omitempty"`
		Name            string             `json:"name,omitempty"`
		Birth           string             `json:"birth,omitempty"`
		Gender          string             `json:"gender,omitempty"`
		Religion        string             `json:"religion,omitempty"`
		Marriage        string             `json:"marriage,omitempty"`
		Hobby           string             `json:"hobby,omitempty"`
		Telephone       string             `json:"telephone,omitempty"`
		Email           string             `json:"email,omitempty"`
		Twitter         string             `json:"twitter,omitempty"`
		Linkedin        string             `json:"linkedin,omitempty"`
		Address         string             `json:"address,omitempty"`
		City            string             `json:"city,omitempty"`
		Province        string             `json:"province,omitempty"`
		State           string             `json:"state,omitempty"`
		About           string             `json:"about,omitempty"`
		Portfolio       string             `json:"portfolio,omitempty"`
		Avatar          *multipart.Form    `json:"avatar,omitempty"`
		KTP             *multipart.Form    `json:"ktp,omitempty"`
		Ijazah          *multipart.Form    `json:"ijazah,omitempty"`
		SKCK            *multipart.Form    `json:"skck,omitempty"`
		CV              *multipart.Form    `json:"cv,omitempty"`
		Certificate     *multipart.Form    `json:"certificate,omitempty"`
		SelfDevelopment SelfDevelopment    `json:"selfdevelopment,omitempty"`
		Education       Education          `json:"lasteducation,omitempty"`
		Career          []Career           `json:"career,omitempty"`
		UpdatedAt       string             `json:"updated_at,omitempty"`
	}
	PersonResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
