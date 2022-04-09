package repository

import (
	"context"
	"golangapi/app/models"
	"golangapi/libs"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	mongoDB *mongo.Database
}

func NewPostRepository(mongo *mongo.Database) PostRepositoryI {
	return &PostRepository{
		mongoDB: mongo,
	}
}

func (pr PostRepository) GetAll(ctx context.Context) (res models.PostResponse, err error) {

	results, err := pr.mongoDB.Collection("posts").Find(ctx, bson.M{})
	if err != nil {
		return models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting posts",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var posts []models.PostEntity
	for results.Next(ctx) {
		var row models.PostEntity
		err := results.Decode(&row)
		if err != nil {
			return models.PostResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting posts",
				Data:    nil,
			}, err
		}
		posts = append(posts, row)
	}

	return models.PostResponse{
		Status:  fiber.StatusOK,
		Message: "Success get posts",
		Data: &fiber.Map{
			"data": posts,
		},
	}, err
}

func (pr PostRepository) GetAllByUser(ctx context.Context, id string) (res models.PostResponse, err error) {

	reqId, _ := primitive.ObjectIDFromHex(id)
	results, err := pr.mongoDB.Collection("posts").Find(ctx, bson.M{"authorid": reqId})
	if err != nil {
		return models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting posts",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var posts []models.PostEntity
	for results.Next(ctx) {
		var row models.PostEntity
		err := results.Decode(&row)
		if err != nil {
			return models.PostResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting posts",
				Data:    nil,
			}, err
		}
		posts = append(posts, row)
	}

	return models.PostResponse{
		Status:  fiber.StatusOK,
		Message: "Success get posts",
		Data: &fiber.Map{
			"data": posts,
		},
	}, err
}

func (pr PostRepository) GetOne(ctx context.Context, id string) (res models.PostResponse, err error) {
	var post models.PostEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := pr.mongoDB.Collection("posts").FindOne(ctx, bson.M{"id": reqId}).Decode(&post); err != nil {
		return models.PostResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting posts",
			Data:    nil,
		}, err
	}

	return models.PostResponse{
		Status:  fiber.StatusOK,
		Message: "Success get posts",
		Data: &fiber.Map{
			"data": post,
		},
	}, err
}

func (pr PostRepository) Insert(ctx context.Context, req models.PostRequest) (res int, err error) {
	metadata, err := libs.MultipleFileHandler(req.Images, req.Id.Hex())

	if metadata[0] == nil {
		return fiber.StatusInternalServerError, err
	}

	images := []models.Image{}

	for i := range metadata[0] {
		newImage := &models.Image{
			Id:   primitive.NewObjectID(),
			Name: metadata[0][i],
			Url:  metadata[1][i],
		}
		images = append(images, *newImage)
	}

	newPost := models.PostEntity{
		Id:         primitive.NewObjectID(),
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		Price:      req.Price,
		AuthorId:   req.AuthorId,
		AuthorName: req.AuthorName,
		Images:     images,
	}

	if _, err := pr.mongoDB.Collection("posts").InsertOne(ctx, newPost); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusCreated, nil
}

func (pr PostRepository) Update(ctx context.Context, req models.PostRequest) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())

	if _, err = pr.mongoDB.Collection("posts").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": req}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated post details
	var updatedpost models.PostEntity
	if err := pr.mongoDB.Collection("posts").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedpost); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (pr PostRepository) Delete(ctx context.Context, id string) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(id)

	if _, err = pr.mongoDB.Collection("posts").DeleteOne(ctx, bson.M{"id": reqId}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}