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

type NewsRepository struct {
	mongoDB *mongo.Database
}

func NewNewsRepository(mongo *mongo.Database) NewsRepositoryI {
	return &NewsRepository{
		mongoDB: mongo,
	}
}

func (ur NewsRepository) InsertNews(ctx context.Context, req models.NewsRequest) (res int, err error) {
	dt := time.Now()

	newNews := models.NewsEntity{
		Id:        primitive.NewObjectID(),
		Author:    req.Author,
		Type:      req.Type,
		Title:     req.Title,
		Body1:     req.Body1,
		Body2:     req.Body2,
		Quotes:    req.Quotes,
		Header:    req.Header,
		CreatedAt: dt.Format("01/02/2006"),
	}

	if _, err := ur.mongoDB.Collection("news").InsertOne(ctx, newNews); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur NewsRepository) GetAllNews(ctx context.Context) (res models.NewsResponse, err error) {

	results, err := ur.mongoDB.Collection("news").Find(ctx, bson.M{})
	if err != nil {
		return models.NewsResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting news",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var news []models.NewsEntity
	for results.Next(ctx) {
		var row models.NewsEntity
		err := results.Decode(&row)
		if err != nil {
			return models.NewsResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting news",
				Data:    nil,
			}, err
		}
		news = append(news, row)
	}

	return models.NewsResponse{
		Status:  fiber.StatusOK,
		Message: "Success get news",
		Data: &fiber.Map{
			"data": news,
		},
	}, err
}

func (ur NewsRepository) GetOneNews(ctx context.Context, id string) (res models.NewsResponse, err error) {
	var news models.NewsEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := ur.mongoDB.Collection("news").FindOne(ctx, bson.M{"id": reqId}).Decode(&news); err != nil {
		return models.NewsResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting news",
			Data:    nil,
		}, err
	}

	return models.NewsResponse{
		Status:  fiber.StatusOK,
		Message: "Success get news",
		Data: &fiber.Map{
			"data": news,
		},
	}, err
}

func (pr NewsRepository) UpdateNews(ctx context.Context, req models.NewsRequest) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	if _, err = pr.mongoDB.Collection("news").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": bson.M{
		"author": req.Author,
		"type":   req.Type,
		"title":  req.Title,
		"body1":  req.Body1,
		"body2":  req.Body2,
		"quotes": req.Quotes,
		"header": req.Header,
	}}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedNews models.NewsEntity
	if err := pr.mongoDB.Collection("news").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedNews); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (pr NewsRepository) DeleteNews(ctx context.Context, id string) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(id)

	if _, err = pr.mongoDB.Collection("news").DeleteOne(ctx, bson.M{"id": reqId}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}
