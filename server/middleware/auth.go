package middleware

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(c *fiber.Ctx) error {
	cookieToken := c.Cookies("authToken")

	if cookieToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Missing JWT",
		})
	}

	token, err := jwt.ParseWithClaims(cookieToken, &jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("LOGIN_KEY")), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Token expired, please log in again",
			})
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse jwt claims",
		})
	}

	email, _ := (*claims)["sub"].(string)

	c.Locals("email", email)
	return c.Next()
}
