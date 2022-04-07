package delivery

import (
	"golangapi/app/delivery/controllers"
	"golangapi/app/usecase"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App, u usecase.PostUsecaseI) {
	c := controllers.NewPostController(app, u)

	app.Get("/api/posts", c.GetAllPost)                    // Get all posts
	app.Get("/api/posts/user/:userId", c.GetAllPostByUser) // Get all posts by user id
	app.Get("/api/post/:postId", c.GetOnePost)             // Get a single post
	app.Post("/api/post", c.InsertPost)                    // Create a new post
	app.Put("/api/post", c.UpdatePost)                     // Update an existing post
	app.Delete("/api/post/:postId", c.DeletePost)          // Delete post
}
