package controllers

import (
	"task-inator3000/config"
	"task-inator3000/models"
	"task-inator3000/utils"

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
			"error": "invalid tasklist id",
		})
	}

	task := new(models.Task)

	err = c.BodyParser(task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
		})
	}

	task.ID = primitive.NewObjectID()

	err = task.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// encrypt content
	task.Content, err = utils.AESEncrypt(task.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
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
			"error": "error while adding task:\n" + err.Error(),
		})
	}

	// if no tasklist updated
	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "taskList not found / permission denied",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"taskID": task.ID,
	})
}

func DeleteTask(c *fiber.Ctx) error {
	taskLists := config.DB.Collection("task_lists")

	// Convert the param taskID from string to ObjectID
	taskID, err := primitive.ObjectIDFromHex(c.Params("taskID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid task id",
		})
	}

	// update the task list by removing the task with the given taskID
	filter := bson.M{
		"tasks._id": taskID,
		"owner":     c.Locals("email"),
	}
	update := bson.M{
		"$pull": bson.M{
			"tasks": bson.M{"_id": taskID},
		},
	}
	result, err := taskLists.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found / permission denied",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
