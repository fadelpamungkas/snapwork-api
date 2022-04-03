package main

import (
	"fmt"
	"golangapi/configs"
	"golangapi/routes"
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
		Prefork:      true,
		ServerHeader: "Fiber",
		AppName:      "Snapwork-API",
	})

	app.Use(cors.New())

	configs.ConnectDB()

	routes.UserRoute(app)
	routes.PostRoute(app)

	// Heroku automatically assigns a port our web server.
	// If it fails we instruct it to use port 000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server started on port " + port)
	app.Listen(":" + port)

}
