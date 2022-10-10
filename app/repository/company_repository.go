package repository

import (
	"context"
	"golangapi/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	mongoDB *mongo.Database
}

func NewCompanyRepository(mongo *mongo.Database) CompanyRepositoryI {
	return &CompanyRepository{
		mongoDB: mongo,
	}
}

func (ur CompanyRepository) InsertCompany(ctx context.Context, req models.CompanyRequest) (res int, err error) {
	dt := time.Now()

	newCompany := models.CompanyEntity{
		Id:            primitive.NewObjectID(),
		UserId:        req.UserId,
		Status:        "Pending",
		Name:          req.Name,
		Email:         req.Email,
		IndustryType:  req.IndustryType,
		Website:       req.Website,
		Phone:         req.Phone,
		Description:   req.Description,
		Country:       req.Country,
		Province:      req.Province,
		City:          req.City,
		Address:       req.Address,
		PostalCode:    req.PostalCode,
		OfficerName:   req.OfficerName,
		OfficerEmail:  req.OfficerEmail,
		OfficerPhone:  req.OfficerPhone,
		OfficerMobile: req.OfficerMobile,
		CreatedAt:     dt.Format("01/02/2006"),
	}

	if _, err := ur.mongoDB.Collection("companydata").InsertOne(ctx, newCompany); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur CompanyRepository) GetAllCompanies(ctx context.Context) (res models.CompanyResponse, err error) {

	results, err := ur.mongoDB.Collection("companydata").Find(ctx, bson.M{})
	if err != nil {
		return models.CompanyResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting company",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var companies []models.CompanyEntity
	for results.Next(ctx) {
		var row models.CompanyEntity
		err := results.Decode(&row)
		if err != nil {
			return models.CompanyResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting companies",
				Data:    nil,
			}, err
		}
		companies = append(companies, row)
	}

	return models.CompanyResponse{
		Status:  fiber.StatusOK,
		Message: "Success get companies",
		Data: &fiber.Map{
			"data": companies,
		},
	}, err
}

func (ur CompanyRepository) InsertJob(ctx context.Context, req models.CompanyJobRequest) (res int, err error) {
	dt := time.Now()

	newJob := models.CompanyJobEntity{
		Id:          primitive.NewObjectID(),
		UserId:      req.UserId,
		Name:        req.Name,
		Kind:        req.Kind,
		Type:        req.Type,
		Status:      req.Status,
		Description: req.Description,
		SoftSkill:   req.SoftSkill,
		HardSkill:   req.HardSkill,
		Education:   req.Education,
		Major:       req.Major,
		SpecificReq: req.SpecificReq,
		Placement:   req.Placement,
		Available:   req.Available,
		CreatedAt:   dt.Format("01/02/2006 15:04:05"),
		UpdatedAt:   dt.Format("01/02/2006 15:04:05"),
	}

	if _, err := ur.mongoDB.Collection("companyjobs").InsertOne(ctx, newJob); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}
//
// func (ur CompanyRepository) GetAllJobsInCompany(ctx context.Context) (res models.CompanyJobResponse, err error) {
//
// 	results, err := ur.mongoDB.Collection("companyjobs").Find(ctx, bson.M{})
// 	if err != nil {
// 		return models.CompanyJobResponse{
// 			Status:  fiber.StatusInternalServerError,
// 			Message: "Error getting company job",
// 			Data:    nil,
// 		}, err
// 	}
// 	defer results.Close(ctx)
//
// 	var jobs []models.CompanyJobEntity
// 	for results.Next(ctx) {
// 		var row models.CompanyJobEntity
// 		err := results.Decode(&row)
// 		if err != nil {
// 			return models.CompanyJobResponse{
// 				Status:  fiber.StatusInternalServerError,
// 				Message: "Error getting company jobs",
// 				Data:    nil,
// 			}, err
// 		}
// 		jobs = append(jobs, row)
// 	}
//
// 	return models.CompanyJobResponse{
// 		Status:  fiber.StatusOK,
// 		Message: "Success get company jobs",
// 		Data: &fiber.Map{
// 			"data": jobs,
// 		},
// 	}, err
// }

func (ur CompanyRepository) GetAllJobsInCompany(ctx context.Context, id string) (res models.CompanyJobResponse, err error) {
	var jobs models.CompanyJobEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := ur.mongoDB.Collection("companyjobs").FindOne(ctx, bson.M{"userid": reqId}).Decode(&jobs); err != nil {
		return models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting company jobs",
			Data:    nil,
		}, err
	}

	return models.CompanyJobResponse{
		Status:  fiber.StatusOK,
		Message: "Success get company jobs",
		Data: &fiber.Map{
			"data": jobs,
		},
	}, err
}

