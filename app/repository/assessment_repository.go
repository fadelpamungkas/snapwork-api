package repository

import (
	"context"
	"golangapi/app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AssessmentRepository struct {
	mongoDB *mongo.Database
}

func NewAssessmentRepository(mongo *mongo.Database) AssessmentRepositoryI {
	return &AssessmentRepository{
		mongoDB: mongo,
	}
}

func (ur AssessmentRepository) InsertAssessment(ctx context.Context, req models.AssessmentRequest) (res int, err error) {

	newAssessment := models.AssessmentEntity{
		Id:     primitive.NewObjectID(),
		Item_1: req.Item_1,
		Item_2: req.Item_2,
		Item_3: req.Item_3,
		Item_4: req.Item_4,
	}

	if _, err := ur.mongoDB.Collection("teskepribadian").InsertOne(ctx, newAssessment); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur AssessmentRepository) GetAllAssessment(ctx context.Context) (res models.AssessmentResponse, err error) {

	results, err := ur.mongoDB.Collection("teskepribadian").Find(ctx, bson.M{})
	if err != nil {
		return models.AssessmentResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting assessment",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var assessments []models.AssessmentEntity
	for results.Next(ctx) {
		var row models.AssessmentEntity
		err := results.Decode(&row)
		if err != nil {
			return models.AssessmentResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting assessment",
				Data:    nil,
			}, err
		}
		assessments = append(assessments, row)
	}

	return models.AssessmentResponse{
		Status:  fiber.StatusOK,
		Message: "Success get assessment",
		Data: &fiber.Map{
			"data": assessments,
		},
	}, err
}
