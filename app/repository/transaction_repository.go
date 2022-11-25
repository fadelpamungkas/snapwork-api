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

type TransactionRepository struct {
	mongoDB *mongo.Database
}

func NewTransactionRepository(mongo *mongo.Database) TransactionRepositoryI {
	return &TransactionRepository{
		mongoDB: mongo,
	}
}

func (ur TransactionRepository) InsertOrder(ctx context.Context, req models.OrderRequest) (res int, err error) {
	dt := time.Now()

	newOrder := models.OrderEntity{
		Id:        primitive.NewObjectID(),
		UserId:    req.UserId,
		Status:    "Pending",
		Name:      req.Name,
		Method:    req.Method,
		Amount:    req.Amount,
		FileProof: req.FileProof,
		CreatedAt: dt.Format("01/02/2006 15:04:05"),
	}

	if _, err := ur.mongoDB.Collection("transaction").InsertOne(ctx, newOrder); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur TransactionRepository) GetAllOrder(ctx context.Context) (res models.OrderResponse, err error) {

	results, err := ur.mongoDB.Collection("transaction").Find(ctx, bson.M{})
	if err != nil {
		return models.OrderResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting order data",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var orders []models.OrderEntity
	for results.Next(ctx) {
		var row models.OrderEntity
		err := results.Decode(&row)
		if err != nil {
			return models.OrderResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting order data",
				Data:    nil,
			}, err
		}
		orders = append(orders, row)
	}

	return models.OrderResponse{
		Status:  fiber.StatusOK,
		Message: "Success get order data",
		Data: &fiber.Map{
			"data": orders,
		},
	}, err
}

func (ur TransactionRepository) InsertApplication(ctx context.Context, req models.ApplicationRequest) (res int, err error) {
	dt := time.Now()

	applicationId := primitive.NewObjectID()
	personId, _ := primitive.ObjectIDFromHex(req.PersonId.Hex())
	companyId, _ := primitive.ObjectIDFromHex(req.CompanyId.Hex())

	applicationsUser := []models.ApplicationUserEntity{}
	applicationsCompany := []models.ApplicationCompanyEntity{}

	//get updated post details
	var currentUserData models.PersonEntity
	if err := ur.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": personId}).Decode(&currentUserData); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var currentCompanyData models.CompanyEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": companyId}).Decode(&currentCompanyData); err != nil {
		return fiber.StatusInternalServerError, err
	}

	newUserApplication := models.ApplicationUserEntity{
		Id:           applicationId,
		CompanyId:    req.CompanyId,
		CompanyJobId: req.CompanyJobId,
		Status:       "Applied",
		CompanyName:  req.CompanyName,
		JobPosition:  req.JobPosition,
		JobPlacement: req.JobPlacement,
		JobType:      req.JobType,
		CreatedAt:    dt.Format("01/02/2006 15:04:05"),
		UpdatedAt:    dt.Format("01/02/2006 15:04:05"),
	}

	newCompanyApplication := models.ApplicationCompanyEntity{
		Id:            applicationId,
		PersonId:      req.PersonId,
		CompanyId:     req.CompanyId,
		CompanyJobId:  req.CompanyJobId,
		Status:        "Applied",
		CompanyName:   req.CompanyName,
		JobPosition:   req.JobPosition,
		JobPlacement:  req.JobPlacement,
		JobType:       req.JobType,
		UserName:      req.UserName,
		UserMarriage:  req.UserMarriage,
		UserState:     req.UserState,
		UserBirth:     req.UserBirth,
		UserGender:    req.UserGender,
		UserAddress:   req.UserAddress,
		UserAbout:     req.UserAbout,
		UserEmail:     req.UserEmail,
		UserTelephone: req.UserTelephone,
		UserTwitter:   req.UserTwitter,
		UserLinkedin:  req.UserLinkedin,
		UserDocument:  req.UserDocument,
		UserEducation: req.UserEducation,
		CreatedAt:     dt.Format("01/02/2006 15:04:05"),
		UpdatedAt:     dt.Format("01/02/2006 15:04:05"),
	}

	applicationsUser = append(currentUserData.Applications, newUserApplication)
	applicationsCompany = append(currentCompanyData.Applications, newCompanyApplication)

	if _, err = ur.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": personId}, bson.M{"$set": bson.M{
		"applications": applicationsUser,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	if _, err = ur.mongoDB.Collection("companydata").UpdateOne(ctx, bson.M{"id": companyId}, bson.M{"$set": bson.M{
		"applications": applicationsCompany,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedUserApplications models.ApplicationUserEntity
	if err := ur.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": personId}).Decode(&updatedUserApplications); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedCompanyApplications models.ApplicationUserEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": companyId}).Decode(&updatedCompanyApplications); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur TransactionRepository) UpdateApplicationStatus(ctx context.Context, req models.ApplicationStatusRequest) (res int, err error) {
	dt := time.Now()

	applicationId, _ := primitive.ObjectIDFromHex(req.ApplicationId.Hex())
	companyId, _ := primitive.ObjectIDFromHex(req.CompanyId.Hex())
	personId, _ := primitive.ObjectIDFromHex(req.PersonId.Hex())
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	var status = [6]string{"Applied", "Screening", "Interview", "Test", "Accepted", "Rejected"}
	var currentStatus string

	switch req.Status {
	case status[0]:
		currentStatus = status[1]
	case status[1]:
		currentStatus = status[2]
	case status[2]:
		currentStatus = status[3]
	case status[3]:
		currentStatus = status[4]
	case status[4]:
		currentStatus = status[4]
	case status[5]:
		currentStatus = status[5]
	default:
		currentStatus = status[0]
	}

	updateData := bson.M{
		"applications.$.status":     currentStatus,
		"applications.$.updated_at": dt.Format("01/02/2006 15:04:05"),
	}

	if _, err = ur.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": personId, "applications.id": applicationId}, bson.M{"$set": updateData}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	if _, err = ur.mongoDB.Collection("companydata").UpdateOne(ctx, bson.M{"id": companyId, "applications.id": applicationId}, bson.M{"$set": updateData}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	var updatedUserApplications models.ApplicationUserEntity
	if err := ur.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": personId}).Decode(&updatedUserApplications); err != nil {
		return fiber.StatusInternalServerError, err
	}

	var updatedCompanyApplications models.ApplicationUserEntity
	if err := ur.mongoDB.Collection("companydata").FindOne(ctx, bson.M{"id": companyId}).Decode(&updatedCompanyApplications); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}
