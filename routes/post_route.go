package routes

import (
	"golangapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App) {
	app.Get("/api/posts", controllers.GetAllPosts)          // Get all posts
	app.Get("/api/post/:postId", controllers.GetSinglePost) // Get a single post
	app.Post("/api/posts", controllers.CreatePost)          // Create a new post
	app.Put("/api/post/:postId", controllers.UpdatePost)    // Update an existing post
	app.Delete("/api/post/:postId", controllers.DeletePost) // Delete post
}
