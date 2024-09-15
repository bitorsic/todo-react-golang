package middleware

import (
	"golang-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func VerifyJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// Check if it starts with "Bearer "
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"err":     "invalid authorization header format",
		})
	}

	authToken := authHeader[7:]

	if authToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"err":     "missing auth token",
		})
	}

	email, err := utils.VerifyJWT(authToken, false)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"err":     err.Error(),
		})
	}

	// cookieToken := c.Cookies("refreshToken")

	// if cookieToken == "" {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"success": false,
	// 		"err":     "missing refresh token",
	// 	})
	// }

	// email, err := utils.VerifyJWT(cookieToken, true)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"success": false,
	// 		"err":     err.Error(),
	// 	})
	// }

	c.Locals("email", email)
	return c.Next()
}
