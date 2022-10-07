package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CompanyEntity struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		UserId        string             `json:"user_id,omitempty"`
		Status        string             `json:"status,omitempty"`
		Name          string             `json:"name,omitempty"`
		Email         string             `json:"email,omitempty"`
		IndustryType  string             `json:"industrytype,omitempty"`
		Website       string             `json:"website,omitempty""`
		Phone         string             `json:"phone,omitempty""`
		Description   string             `json:"description,omitempty""`
		Country       string             `json:"country,omitempty""`
		Province      string             `json:"province,omitempty""`
		City          string             `json:"city,omitempty""`
		Address       string             `json:"address,omitempty""`
		OfficerName   string             `json:"officername,omitempty""`
		OfficerEmail  string             `json:"officeremail,omitempty""`
		OfficerPhone  string             `json:"officerphone,omitempty""`
		OfficerMobile string             `json:"officermobile,omitempty""`
		CreatedAt     string             `json:"created_at,omitempty""`
	}
	CompanyRequest struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		UserId        string             `json:"user_id,omitempty"`
		Status        string             `json:"status,omitempty"`
		Name          string             `json:"name,omitempty"`
		Email         string             `json:"email,omitempty"`
		IndustryType  string             `json:"industrytype,omitempty"`
		Website       string             `json:"website,omitempty""`
		Phone         string             `json:"phone,omitempty""`
		Description   string             `json:"description,omitempty""`
		Country       string             `json:"country,omitempty""`
		Province      string             `json:"province,omitempty""`
		City          string             `json:"city,omitempty""`
		Address       string             `json:"address,omitempty""`
		OfficerName   string             `json:"officername,omitempty""`
		OfficerEmail  string             `json:"officeremail,omitempty""`
		OfficerPhone  string             `json:"officerphone,omitempty""`
		OfficerMobile string             `json:"officermobile,omitempty""`
		CreatedAt     string             `json:"created_at,omitempty""`
	}
	CompanyResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
