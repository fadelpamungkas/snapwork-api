package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ApplicationCompanyEntity struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		PersonId      primitive.ObjectID `json:"personid,omitempty"`
		CompanyId     primitive.ObjectID `json:"companyid,omitempty"`
		CompanyJobId  primitive.ObjectID `json:"companyjobid,omitempty"`
		Status        string             `json:"status,omitempty"`
		CompanyName   string             `json:"companyname,omitempty"`
		JobPosition   string             `json:"jobposition,omitempty"`
		JobPlacement  string             `json:"jobplacement,omitempty"`
		JobType       string             `json:"jobtype,omitempty"`
		UserName      string             `json:"username,omitempty"`
		UserMarriage  string             `json:"usermarriage,omitempty"`
		UserState     string             `json:"userstate,omitempty"`
		UserBirth     string             `json:"userbirth,omitempty"`
		UserGender    string             `json:"usergender,omitempty"`
		UserAddress   string             `json:"useraddress,omitempty"`
		UserAbout     string             `json:"userabout,omitempty"`
		UserEmail     string             `json:"useremail,omitempty"`
		UserTelephone string             `json:"usertelephone,omitempty"`
		UserTwitter   string             `json:"usertwitter,omitempty"`
		UserLinkedin  string             `json:"userlinkedin,omitempty"`
		UserDocument  Document           `json:"userdocument,omitempty"`
		UserEducation Education          `json:"usereducation,omitempty"`
		CreatedAt     string             `json:"created_at,omitempty"`
		UpdatedAt     string             `json:"updated_at,omitempty"`
	}
	ApplicationUserEntity struct {
		Id           primitive.ObjectID `json:"_id,omitempty"`
		CompanyId    primitive.ObjectID `json:"companyid,omitempty"`
		CompanyJobId primitive.ObjectID `json:"companyjobid,omitempty"`
		Status       string             `json:"status,omitempty"`
		CompanyName  string             `json:"companyname,omitempty"`
		JobPosition  string             `json:"jobposition,omitempty"`
		JobPlacement string             `json:"jobplacement,omitempty"`
		JobType      string             `json:"jobtype,omitempty"`
		CreatedAt    string             `json:"created_at,omitempty"`
		UpdatedAt    string             `json:"updated_at,omitempty"`
	}
	ApplicationRequest struct {
		Id            primitive.ObjectID `json:"_id,omitempty"`
		PersonId      primitive.ObjectID `json:"personid,omitempty"`
		CompanyId     primitive.ObjectID `json:"companyid,omitempty"`
		CompanyJobId  primitive.ObjectID `json:"companyjobid,omitempty"`
		CompanyName   string             `json:"companyname,omitempty"`
		JobPosition   string             `json:"jobposition,omitempty"`
		JobPlacement  string             `json:"jobplacement,omitempty"`
		JobType       string             `json:"jobtype,omitempty"`
		UserName      string             `json:"username,omitempty"`
		UserMarriage  string             `json:"usermarriage,omitempty"`
		UserState     string             `json:"userstate,omitempty"`
		UserBirth     string             `json:"userbirth,omitempty"`
		UserGender    string             `json:"usergender,omitempty"`
		UserAddress   string             `json:"useraddress,omitempty"`
		UserAbout     string             `json:"userabout,omitempty"`
		UserEmail     string             `json:"useremail,omitempty"`
		UserTelephone string             `json:"usertelephone,omitempty"`
		UserTwitter   string             `json:"usertwitter,omitempty"`
		UserLinkedin  string             `json:"userlinkedin,omitempty"`
		UserDocument  Document           `json:"userdocument,omitempty"`
		UserEducation Education          `json:"usereducation,omitempty"`
		CreatedAt     string             `json:"created_at,omitempty"`
		UpdatedAt     string             `json:"updated_at,omitempty"`
	}
	ApplicationStatusRequest struct {
		ApplicationId primitive.ObjectID `json:"applicationid,omitempty"`
		PersonId      primitive.ObjectID `json:"personid,omitempty"`
		CompanyId     primitive.ObjectID `json:"companyid,omitempty"`
		CompanyJobId  primitive.ObjectID `json:"companyjobid,omitempty"`
		Status        string             `json:"status,omitempty"`
	}
	ApplicationResponse struct {
		Status  int        `json:"status"`
		Message string     `json:"message"`
		Data    *fiber.Map `json:"data"`
	}
)
