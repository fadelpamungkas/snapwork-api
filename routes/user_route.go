package routes

import (
	"golangapi/configs"
	"golangapi/controllers"
	"golangapi/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Static("/public", configs.ProjectRootPath+"/public/assets", fiber.Static{
		Browse:   true,
		Compress: false,
	})

	app.Get("/api/users", middleware.Auth, controllers.GetAllUsers)
	app.Get("/api/user/:userId", controllers.GetSingleUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Put("/api/user/:userId", controllers.UpdateUser)
	app.Delete("/api/user/:userId", controllers.DeleteUser)

	app.Get("/login", controllers.Login)
}
