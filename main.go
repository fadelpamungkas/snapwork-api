package main

import (
	"fmt"
	delivery "golangapi/app/delivery/api"
	"golangapi/app/repository"
	"golangapi/app/usecase"
	"golangapi/libs"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	fmt.Println(err)
		// 	return c.Status(500).SendString("Internal Server Error")
		// },
		Prefork:      false,
		ServerHeader: "Fiber",
		AppName:      "Snapwork-API",
	})
	app.Use(cors.New())

	env, err := libs.Environment()

	// Mongo
	mongo, err := libs.Connect(env.Database)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connected to MongoDB")

	// User
	userRepo := repository.NewUserRepository(mongo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	delivery.UserRoute(app, userUsecase)

	// Post
	postRepo := repository.NewPostRepository(mongo)
	postUsecase := usecase.NewPostUsecase(postRepo)
	delivery.PostRoute(app, postUsecase)

	// Post
	transRepo := repository.NewTransactionRepository(mongo)
	transUsecase := usecase.NewTransactionUsecase(transRepo)
	delivery.TransactionRoute(app, transUsecase)

	// Post
	companyRepo := repository.NewCompanyRepository(mongo)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	delivery.CompanyRoute(app, companyUsecase)

	// News
	newsRepo := repository.NewNewsRepository(mongo)
	newsUsecase := usecase.NewNewsUsecase(newsRepo)
	delivery.NewsRoute(app, newsUsecase)

	// Heroku automatically assigns a port our web server.
	// If it fails we instruct it to use port 000
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%v", env.Port)
	}

	log.Println("Server started on port " + port)
	log.Fatal(app.Listen(":" + port))
}
