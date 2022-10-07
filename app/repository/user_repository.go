package repository

import (
	"context"
	"golangapi/app/models"
	"golangapi/libs"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoDB *mongo.Database
}

func NewUserRepository(mongo *mongo.Database) UserRepositoryI {
	return &UserRepository{
		mongoDB: mongo,
	}
}

func (ur UserRepository) Login(ctx context.Context, req models.LoginRequest) (res models.UserResponse, err error) {
	var user models.UserEntity

	if err := ur.mongoDB.Collection("users").FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err != nil {
		return models.UserResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Wrong Credential",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}, err
	}

	// validating password
	if err := libs.CheckHashPassword(req.Password, user.Password); err != nil {
		return models.UserResponse{
			Status:  fiber.StatusUnauthorized,
			Message: "Wrong Credential",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		}, err
	}

	// generate jwt token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id.Hex()
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerate := libs.GenerateToken(&claims)

	if errGenerate != nil {
		return models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error generating token",
			Data: &fiber.Map{
				"data": errGenerate.Error(),
			},
		}, err
	}

	return models.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Login Success",
		Data: &fiber.Map{
			"data": token,
		},
	}, nil
}

func (ur UserRepository) GetAll(ctx context.Context) (res models.UserResponse, err error) {

	results, err := ur.mongoDB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting users",
			Data:    nil,
		}, err
	}
	defer results.Close(ctx)

	var users []models.UserEntity
	for results.Next(ctx) {
		var row models.UserEntity
		err := results.Decode(&row)
		if err != nil {
			return models.UserResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Error getting users",
				Data:    nil,
			}, err
		}
		users = append(users, row)
	}

	return models.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Success get users",
		Data: &fiber.Map{
			"data": users,
		},
	}, err
}

func (ur UserRepository) GetOne(ctx context.Context, id string) (res models.UserResponse, err error) {
	var user models.UserEntity

	reqId, _ := primitive.ObjectIDFromHex(id)
	if err := ur.mongoDB.Collection("users").FindOne(ctx, bson.M{"id": reqId}).Decode(&user); err != nil {
		return models.UserResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Error getting users",
			Data:    nil,
		}, err
	}

	return models.UserResponse{
		Status:  fiber.StatusOK,
		Message: "Success get users",
		Data: &fiber.Map{
			"data": user,
		},
	}, err
}

func (ur UserRepository) Insert(ctx context.Context, req models.UserRequest) (res int, err error) {
	hashPassword, err := libs.HashPassword(req.Password)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	newUser := models.UserEntity{
		Id:       primitive.NewObjectID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
		Role:     "user",
	}

	if _, err := ur.mongoDB.Collection("users").InsertOne(ctx, newUser); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur UserRepository) Update(ctx context.Context, req models.UserRequest) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(req.Id.Hex())
	hashPassword, err := libs.HashPassword(req.Password)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	updateData := bson.M{
		"name":     req.Name,
		"email":    req.Email,
		"password": hashPassword,
	}

	if _, err = ur.mongoDB.Collection("users").UpdateOne(ctx, bson.M{"id": reqId}, bson.M{"$set": updateData}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	//get updated user details
	var updatedUser models.UserEntity
	if err := ur.mongoDB.Collection("users").FindOne(ctx, bson.M{"id": reqId}).Decode(&updatedUser); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func (ur UserRepository) Delete(ctx context.Context, id string) (res int, err error) {
	reqId, _ := primitive.ObjectIDFromHex(id)

	if _, err = ur.mongoDB.Collection("product").DeleteOne(ctx, bson.M{"id": reqId}); err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}
