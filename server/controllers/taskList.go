package controllers

import (
	"golang-backend/config"
	"golang-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTaskLists(c *fiber.Ctx) error {
	taskLists := config.DB.Collection("task_lists")
	email := c.Locals("email").(string)

	// find all task lists for the logged in user
	filter := bson.M{"owner": email}
	opts := options.Find().SetProjection(bson.M{
		"title": 1,
		"tasks": 1,
	})
	cursor, err := taskLists.Find(c.Context(), filter, opts)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while finding in the database",
			"err":     err,
		})
	}

	var result []models.TaskList
	err = cursor.All(c.Context(), &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while iterating cursor",
			"err":     err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":    true,
		"task_lists": result,
	})
}
