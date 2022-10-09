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
