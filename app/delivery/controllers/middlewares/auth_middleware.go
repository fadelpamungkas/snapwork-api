package middlewares

import (
	"fmt"
	"golangapi/libs"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims, err := libs.DecodeToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated token",
			"error":   err.Error(),
		})
	}

	// role := claims["role"].(string)

	// if role != "admin" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "Permission Denied",
	// 	})
	// }

	fmt.Println("authUser", claims)
	c.Locals("authUser", claims)

	return c.Next()
}
