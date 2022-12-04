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

func (ur CompanyRepository) UpdateCompanyStatus(ctx context.Context, req models.CompanyStatusRequest) (res int, err error) {

	companyId, _ := primitive.ObjectIDFromHex(req.CompanyId.Hex())
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	var status string

	switch req.Status {
	case 1:
		status = "Verified"
	case 2:
		status = "Decline"
	default:
		status = "Pending"
	}

	updateData := bson.M{
		"status": status,
	}

	if _, err = ur.mongoDB.Collection("companydata").UpdateOne(ctx, bson.M{"id": companyId}, bson.M{"$set": updateData}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	var updatedCompany models.CompanyEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": companyId}).Decode(&updatedCompany); err != nil {
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

func (ur CompanyRepository) InsertJob(ctx context.Context, req models.CompanyJobRequest) (res models.CompanyJobResponse, err error) {
	dt := time.Now()

	reqId, _ := primitive.ObjectIDFromHex(req.CompanyId.Hex())
	JobId := primitive.NewObjectID()

	jobs := []models.CompanyJobEntity{}

	//get updated post details
	var currentData models.CompanyEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": reqId}).Decode(&currentData); err != nil {
		return models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error find company job",
			Data:    nil,
		}, err
	}

	newJob := models.CompanyJobEntity{
		Id:          JobId,
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
		return models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error insert job",
			Data:    nil,
		}, err
	}

	//get updated post details
	var updatedCompanyJob models.CompanyEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedCompanyJob); err != nil {
		return models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error insert job",
			Data:    nil,
		}, err
	}

	log.Println(updatedCompanyJob)

	return models.CompanyJobResponse{
		Status:  fiber.StatusOK,
		Message: "Success insert job",
		Data: &fiber.Map{
			"data": JobId,
		},
	}, err
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

func (ur CompanyRepository) GetCompanyByUserId(ctx context.Context, id string) (res models.CompanyResponse, err error) {
	var company models.CompanyEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"userid": reqId}).Decode(&company); err != nil {
		return models.CompanyResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting company by userid",
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

func (ur CompanyRepository) GetJobCompany(ctx context.Context, companyId string, jobId string) (res models.CompanyJobResponse, err error) {
	var company models.CompanyEntity
	var job *models.CompanyJobEntity

	reqCompanyId, _ := primitive.ObjectIDFromHex(companyId)
	reqJobId, _ := primitive.ObjectIDFromHex(jobId)
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": reqCompanyId}).Decode(&company); err != nil {
		return models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting company",
			Data:    nil,
		}, err
	}

	for i := range company.CompanyJob {
		if company.CompanyJob[i].Id == reqJobId {
			job = &company.CompanyJob[i]
		}
	}

	if job == nil {
		return models.CompanyJobResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting job",
			Data:    nil,
		}, err

	}

	return models.CompanyJobResponse{
		Status:  fiber.StatusOK,
		Message: "Success get job",
		Data: &fiber.Map{
			"data": job,
		},
	}, err
}

func (pr CompanyRepository) UpdateJobCompany(ctx context.Context, req models.CompanyJobRequest) (res int, err error) {
	companyId, _ := primitive.ObjectIDFromHex(req.CompanyId.Hex())
	jobId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	dt := time.Now()

	if _, err = pr.mongoDB.Collection("companydata").UpdateOne(ctx, bson.M{"id": companyId, "companyjob.id": jobId}, bson.M{"$set": bson.M{
		"companyjob.$.name":        req.Name,
		"companyjob.$.kind":        req.Kind,
		"companyjob.$.type":        req.Type,
		"companyjob.$.status":      req.Status,
		"companyjob.$.description": req.Description,
		"companyjob.$.softskill":   req.SoftSkill,
		"companyjob.$.hardskill":   req.HardSkill,
		"companyjob.$.education":   req.Education,
		"companyjob.$.major":       req.Major,
		"companyjob.$.specificreq": req.SpecificReq,
		"companyjob.$.placement":   req.Placement,
		"companyjob.$.available":   req.Available,
		"companyjob.$.updated_at":  dt.Format("01/02/2006 15:04:05"),
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedJobCompany models.CompanyJobEntity
	if err := pr.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": companyId, "companyjob.id": jobId}).Decode(&updatedJobCompany); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (pr CompanyRepository) DeleteJobCompany(ctx context.Context, reqCompanyId string, reqJobId string) (res int, err error) {
	companyId, _ := primitive.ObjectIDFromHex(reqCompanyId)
	jobId, _ := primitive.ObjectIDFromHex(reqJobId)

	if _, err = pr.mongoDB.Collection("companydata").UpdateOne(ctx, bson.M{"id": companyId}, bson.M{"$pull": bson.M{"companyjob": bson.M{"id": jobId}}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}
