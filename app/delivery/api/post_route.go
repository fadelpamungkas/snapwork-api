package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App, u usecase.PostUsecaseI) {
	c := controllers.NewPostController(app, u)

	api := app.Group("/api")

	api.Get("/posts", c.GetAllPost)                    // Get all posts
	api.Get("/posts/user/:userId", c.GetAllPostByUser) // Get all posts by user id
	api.Get("/post/:postId", c.GetOnePost)             // Get a single post
	api.Post("/post", c.InsertPost)                    // Create a new post
	api.Put("/post", c.UpdatePost)                     // Update an existing post
	api.Delete("/post/:postId", c.DeletePost)          // Delete post
}
