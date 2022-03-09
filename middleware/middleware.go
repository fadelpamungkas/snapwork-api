package middleware

import (
	"golangapi/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	if _, err := utils.VerifyToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated token",
			"error":   err.Error(),
		})
	}

	return c.Next()

}
