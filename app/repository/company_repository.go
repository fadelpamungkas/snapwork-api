package repository

import (
	"context"
	"golangapi/app/models"
	"log"
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

	reqId, _ := primitive.ObjectIDFromHex(req.CompanyId.Hex())

	jobs := []models.CompanyJobEntity{}

	//get updated post details
	var currentData models.CompanyEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": reqId}).Decode(&currentData); err != nil {
		return fiber.StatusInternalServerError, err
	}

	newJob := models.CompanyJobEntity{
		Id:          primitive.NewObjectID(),
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

	jobs = append(currentData.CompanyJob, newJob)

	if _, err = ur.mongoDB.Collection("companydata").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"companyjob": jobs,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedCompanyJob models.CompanyEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedCompanyJob); err != nil {
		return fiber.StatusInternalServerError, err
	}

	log.Println(updatedCompanyJob)

	return fiber.StatusOK, nil
}

func (ur CompanyRepository) GetCompany(ctx context.Context, id string) (res models.CompanyResponse, err error) {
	var company models.CompanyEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": reqId}).Decode(&company); err != nil {
		return models.CompanyResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting company",
			Data:    nil,
		}, err
	}

	return models.CompanyResponse{
		Status:  fiber.StatusOK,
		Message: "Success get company",
		Data: &fiber.Map{
			"data": company,
		},
	}, err
}
