package controllers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	return c.SendString("Welcome, " + email)
}
