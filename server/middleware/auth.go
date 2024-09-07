package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(c *fiber.Ctx) error {
	if c.Get("Authorization") == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	authHeader := strings.Split(c.Get("Authorization"), " ")[1]

	token, err := jwt.ParseWithClaims(authHeader, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("LOGIN_KEY")), nil
	})
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse jwt claims",
		})
	}

	email, _ := (*claims)["sub"].(string)

	c.Locals("email", email)
	return c.Next()
}
