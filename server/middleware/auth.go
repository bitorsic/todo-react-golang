package middleware

import (
	"golang-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func VerifyAuthToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// Check if it starts with "Bearer "
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid authorization header format",
		})
	}

	authToken := authHeader[7:]

	if authToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing auth token",
		})
	}

	email, err := utils.VerifyJWT(authToken, false)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Locals("email", email)
	return c.Next()
}
