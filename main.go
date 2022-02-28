package main

import (
	"golangapi/configs"
	"golangapi/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	//Heroku automatically assigns a port our web server. If it
	//fails we instruct it to use port 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(port)
}
