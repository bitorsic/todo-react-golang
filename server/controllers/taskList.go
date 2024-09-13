package controllers

import (
	"golang-backend/config"
	"golang-backend/models"
	"golang-backend/utils"

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
			"err":     err.Error(),
		})
	}

	var result []models.TaskList
	err = cursor.All(c.Context(), &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while iterating cursor",
			"err":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":    true,
		"task_lists": result,
	})
}

func AddTaskList(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	type Input struct {
		Title string `json:"title"`
	}

	var input Input

	// parsing req body to user
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	taskList, err := utils.CreateTaskList(email, input.Title, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating TaskList",
			"err":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":   true,
		"task_list": *taskList,
	})
}
