package controllers

import (
	"golang-backend/config"
	"golang-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTask(c *fiber.Ctx) error {
	taskLists := config.DB.Collection("task_lists")

	// Convert the param taskListID from string to ObjectID
	taskListID, err := primitive.ObjectIDFromHex(c.Params("taskListID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Task List ID",
		})
	}

	task := new(models.Task)

	err = c.BodyParser(task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	task.ID = primitive.NewObjectID()

	err = task.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// appending the task to the task_lists field of the user
	options := bson.M{
		"$push": bson.M{
			"tasks": task,
		},
	}
	_, err = taskLists.UpdateByID(c.Context(), taskListID, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while adding task",
			"err":     err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Task Added",
		"taskID":  task.ID,
	})
}
