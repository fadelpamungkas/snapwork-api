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

type PersonRepository struct {
	mongoDB *mongo.Database
}

func NewPersonRepository(mongo *mongo.Database) PersonRepositoryI {
	return &PersonRepository{
		mongoDB: mongo,
	}
}

func (ur PersonRepository) InsertPerson(ctx context.Context, req models.PersonRequest) (res int, err error) {
	dt := time.Now()

	newPerson := models.PersonEntity{
		Id:        primitive.NewObjectID(),
		UserId:    req.UserId,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: dt.Format("01/02/2006"),
	}

	if _, err := ur.mongoDB.Collection("persondata").InsertOne(ctx, newPerson); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur PersonRepository) GetAllPerson(ctx context.Context) (res models.PersonResponse, err error) {

	results, err := ur.mongoDB.Collection("persondata").Find(ctx, bson.M{})
	if err != nil {
		return models.PersonResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting person",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var persons []models.PersonEntity
	for results.Next(ctx) {
		var row models.PersonEntity
		err := results.Decode(&row)
		if err != nil {
			return models.PersonResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting companies",
				Data:    nil,
			}, err
		}
		persons = append(persons, row)
	}

	return models.PersonResponse{
		Status:  fiber.StatusOK,
		Message: "Success get companies",
		Data: &fiber.Map{
			"data": persons,
		},
	}, err
}

func (ur PersonRepository) GetPerson(ctx context.Context, id string) (res models.PersonResponse, err error) {
	var person models.PersonEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := ur.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"userid": reqId}).Decode(&person); err != nil {
		return models.PersonResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting person data",
			Data:    nil,
		}, err
	}

	return models.PersonResponse{
		Status:  fiber.StatusOK,
		Message: "Success get person data",
		Data: &fiber.Map{
			"data": person,
		},
	}, err
}

func (pr PersonRepository) UpdatePerson(ctx context.Context, req models.PersonRequest) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	dt := time.Now()

	if _, err = pr.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"name":                 req.Name,
		"birth":                req.Birth,
		"gender":               req.Gender,
		"religion":             req.Religion,
		"marriage":             req.Marriage,
		"hobby":                req.Hobby,
		"telephone":            req.Telephone,
		"email":                req.Email,
		"twitter":              req.Twitter,
		"linkedin":             req.Linkedin,
		"address":              req.Address,
		"city":                 req.City,
		"province":             req.Province,
		"state":                req.State,
		"about":                req.About,
		"education.s1":         req.Education.S1,
		"education.s1major":    req.Education.S1Major,
		"education.s1datein":   req.Education.S1DateIn,
		"education.s1dateout":  req.Education.S1DateOut,
		"education.sma":        req.Education.SMA,
		"education.smamajor":   req.Education.SMAMajor,
		"education.smadatein":  req.Education.SMADateIn,
		"education.smadateout": req.Education.SMADateOut,
		"updated_at":           dt.Format("01/02/2006"),
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedPerson models.PersonEntity
	if err := pr.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedPerson); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur PersonRepository) InsertNotification(ctx context.Context, req models.Notification) (res int, err error) {
	dt := time.Now()

	reqId, _ := primitive.ObjectIDFromHex(req.UserId.Hex())

	notifications := []models.Notification{}

	//get updated post details
	var currentData models.PersonEntity
	if err := ur.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"userid": reqId}).Decode(&currentData); err != nil {
		return fiber.StatusInternalServerError, err
	}

	newNotification := models.Notification{
		Id:          primitive.NewObjectID(),
		UserId:      req.UserId,
		Status:      req.Status,
		Title:       req.Title,
		Description: req.Description,
		IsRead:      req.IsRead,
		CreatedAt:   dt.Format("01/02/2006 15:04:05"),
	}

	notifications = append(currentData.Notification, newNotification)

	if _, err = ur.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"userid": reqId}, bson.M{"$set": bson.M{
		"notification": notifications,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedNotification models.Notification
	if err := ur.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"userid": reqId}).Decode(&updatedNotification); err != nil {
		return fiber.StatusInternalServerError, err
	}

	log.Println(updatedNotification)

	return fiber.StatusOK, nil
}

func (pr PersonRepository) UpdateDocumentPerson(ctx context.Context, req models.PersonDocumentRequest) (res int, err error) {

	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	if _, err = pr.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"document.avatar":          req.Avatar,
		"document.avatarname":      req.AvatarName,
		"document.ktp":             req.KTP,
		"document.ktpname":         req.KTPName,
		"document.ijazah":          req.Ijazah,
		"document.ijazahname":      req.IjazahName,
		"document.skck":            req.SKCK,
		"document.skckname":        req.SKCKName,
		"document.cv":              req.CV,
		"document.cvname":          req.CVName,
		"document.certificate":     req.Certificate,
		"document.certificatename": req.CertificateName,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedPerson models.PersonEntity
	if err := pr.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedPerson); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (pr PersonRepository) UpdatePortfolioPerson(ctx context.Context, req models.Portfolio) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	if _, err = pr.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"portfolio.link":        req.Link,
		"portfolio.description": req.Description,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedPerson models.PersonEntity
	if err := pr.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedPerson); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (pr PersonRepository) UpdateSelfDevelopmentPerson(ctx context.Context, req models.SelfDevelopment) (res int, err error) {

	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	if _, err = pr.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"selfdevelopment.score": req.Score,
		"selfdevelopment.file":  req.File,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedPerson models.PersonEntity
	if err := pr.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedPerson); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (pr PersonRepository) UpdateSelfDevelopmentPaymentPerson(ctx context.Context, req models.SelfDevelopmentPaymentRequest) (res int, err error) {

	dt := time.Now()

	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	if _, err = pr.mongoDB.Collection("persondata").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"selfdevelopment.payment.status":    req.Status,
		"selfdevelopment.payment.price":     req.Price,
		"selfdevelopment.payment.fileproof": req.FileProof,
		"selfdevelopment.payment.createdat": dt.Format("01/02/2006 15:04:05"),
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	var updatedPerson models.PersonEntity
	if err := pr.mongoDB.Collection("persondata").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedPerson); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}
