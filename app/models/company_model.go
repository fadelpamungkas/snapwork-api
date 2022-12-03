package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CompanyEntity struct {
		Id            primitive.ObjectID         `json:"_id,omitempty"`
		UserId        primitive.ObjectID         `json:"userid,omitempty"`
		Status        string                     `json:"status,omitempty"`
		Name          string                     `json:"name,omitempty"`
		Email         string                     `json:"email,omitempty"`
		IndustryType  string                     `json:"industrytype,omitempty"`
		Website       string                     `json:"website,omitempty"`
		Phone         string                     `json:"phone,omitempty"`
		Description   string                     `json:"description,omitempty"`
		Country       string                     `json:"country,omitempty"`
		Province      string                     `json:"province,omitempty"`
		City          string                     `json:"city,omitempty"`
		Address       string                     `json:"address,omitempty"`
		PostalCode    string                     `json:"postalcode,omitempty"`
		OfficerName   string                     `json:"officername,omitempty"`
		OfficerEmail  string                     `json:"officeremail,omitempty"`
		OfficerPhone  string                     `json:"officerphone,omitempty"`
		OfficerMobile string                     `json:"officermobile,omitempty"`
		CompanyJob    []CompanyJobEntity         `json:"companyjob,omitempty"`
		Applications  []ApplicationCompanyEntity `json:"applications,omitempty"`
		CreatedAt     string                     `json:"created_at,omitempty"`
	}
	CompanyRequest struct {
		Id            primitive.ObjectID         `json:"_id,omitempty"`
		UserId        primitive.ObjectID         `json:"userid,omitempty"`
		Status        string                     `json:"status,omitempty"`
		Name          string                     `json:"name,omitempty"`
		Email         string                     `json:"email,omitempty"`
		IndustryType  string                     `json:"industrytype,omitempty"`
		Website       string                     `json:"website,omitempty"`
		Phone         string                     `json:"phone,omitempty"`
		Description   string                     `json:"description,omitempty"`
		Country       string                     `json:"country,omitempty"`
		Province      string                     `json:"province,omitempty"`
		City          string                     `json:"city,omitempty"`
		Address       string                     `json:"address,omitempty"`
		PostalCode    string                     `json:"postalcode,omitempty"`
		OfficerName   string                     `json:"officername,omitempty"`
		OfficerEmail  string                     `json:"officeremail,omitempty"`
		OfficerPhone  string                     `json:"officerphone,omitempty"`
		OfficerMobile string                     `json:"officermobile,omitempty"`
		CompanyJob    []CompanyJobEntity         `json:"companyjob,omitempty"`
		Applications  []ApplicationCompanyEntity `json:"applications,omitempty"`
		CreatedAt     string                     `json:"created_at,omitempty"`
	}
	CompanyStatusRequest struct {
		CompanyId primitive.ObjectID `json:"companyid,omitempty"`
		Status    string             `json:"status,omitempty"`
	}
	CompanyResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
	CompanyJobEntity struct {
		Id          primitive.ObjectID `json:"_id,omitempty"`
		Name        string             `json:"name,omitempty"`
		Kind        string             `json:"kind,omitempty"`
		Type        string             `json:"type,omitempty"`
		Status      string             `json:"status,omitempty"`
		Description string             `json:"description,omitempty"`
		SoftSkill   string             `json:"softskill,omitempty"`
		HardSkill   string             `json:"hardskill,omitempty"`
		Education   string             `json:"education,omitempty"`
		Major       string             `json:"major,omitempty"`
		SpecificReq string             `json:"specificreq,omitempty"`
		Placement   string             `json:"placement,omitempty"`
		Available   string             `json:"available,omitempty"`
		CreatedAt   string             `json:"created_at,omitempty"`
		UpdatedAt   string             `json:"updated_at,omitempty"`
	}
	CompanyJobRequest struct {
		Id          primitive.ObjectID `json:"_id,omitempty"`
		CompanyId   primitive.ObjectID `json:"companyid,omitempty"`
		Name        string             `json:"name,omitempty"`
		Kind        string             `json:"kind,omitempty"`
		Type        string             `json:"type,omitempty"`
		Status      string             `json:"status,omitempty"`
		Description string             `json:"description,omitempty"`
		SoftSkill   string             `json:"softskill,omitempty"`
		HardSkill   string             `json:"hardskill,omitempty"`
		Education   string             `json:"education,omitempty"`
		Major       string             `json:"major,omitempty"`
		SpecificReq string             `json:"specificreq,omitempty"`
		Placement   string             `json:"placement,omitempty"`
		Available   string             `json:"available,omitempty"`
		CreatedAt   string             `json:"created_at,omitempty"`
		UpdatedAt   string             `json:"updated_at,omitempty"`
	}
	CompanyJobResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
