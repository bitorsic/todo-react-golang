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

	// appending the task to the tasks field of the tasklist
	filter := bson.M{
		"_id":   taskListID,
		"owner": c.Locals("email"),
	}
	update := bson.M{
		"$push": bson.M{
			"tasks": task,
		},
	}
	result, err := taskLists.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while adding task",
			"err":     err.Error(),
		})
	}

	// if no tasklist updated
	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "TaskList not found or permission denied",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Task Added",
		"taskID":  task.ID,
	})
}
