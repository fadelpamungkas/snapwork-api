package controllers

import (
	"context"
	"golangapi/configs"
	"golangapi/models/entities"
	"golangapi/models/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
var validatePost = validator.New()

func CreatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var post entities.Post
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	//use validator library to validate required fields
	if validationErr := validatePost.Struct(&post); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": validationErr.Error(),
			},
		})
	}

	post.Id = primitive.NewObjectID()

	result, err := postCollection.InsertOne(ctx, post)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error creating post",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{
		Status:  http.StatusCreated,
		Message: "Post created successfully",
		Data: &fiber.Map{
			"data": result,
		},
	})
}

func GetSinglePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postId := c.Params("postId")
	var post entities.Post
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&post)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error getting post",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Post retrieved successfully",
		Data: &fiber.Map{
			"data": post,
		},
	})
}

func UpdatePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postId := c.Params("postId")
	var post entities.Post
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	//validate request body
	if err := c.BodyParser(&post); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request body",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	// update := bson.M{
	// 	"title":      post.Title,
	// 	"content":    post.Content,
	// 	"category":   post.Category,
	// 	"price":      post.Price,
	// 	"authorName": post.AuthorName,
	// }

	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": post})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error updating post",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	//get updated post details
	var updatedPost entities.Post
	if result.MatchedCount == 1 {
		err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedPost)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error getting updated post",
				Data: &fiber.Map{
					"data": err.Error(),
				},
			})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Post updated successfully",
		Data: &fiber.Map{
			"data": updatedPost,
		},
	})
}

func DeletePost(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	postId := c.Params("postId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error deleting post",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{
			Status:  http.StatusNotFound,
			Message: "Post not found",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Post deleted successfully",
		Data: &fiber.Map{
			"data": objId,
		},
	})
}

func GetAllPosts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var posts []entities.Post
	defer cancel()

	results, err := postCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error getting posts",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	//reading from db in an optimal way would
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePost entities.Post
		if err = results.Decode(&singlePost); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error getting posts",
				Data: &fiber.Map{
					"data": err.Error(),
				},
			})
		}
		posts = append(posts, singlePost)
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Posts retrieved successfully",
		Data: &fiber.Map{
			"data": posts,
		},
	})
}
