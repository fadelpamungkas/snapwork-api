package routes

import (
	"golangapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/api/users", controllers.GetAllUsers)
	app.Get("/api/user/:userId", controllers.GetSingleUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Put("/api/user/:userId", controllers.UpdateUser)
	app.Delete("/api/user/:userId", controllers.DeleteUser)
}
