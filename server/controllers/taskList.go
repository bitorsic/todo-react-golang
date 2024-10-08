package controllers

import (
	"task-inator3000/config"
	"task-inator3000/models"
	"task-inator3000/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			"error": "error while finding in the database:\n" + err.Error(),
		})
	}

	var result []models.TaskList
	err = cursor.All(c.Context(), &result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while iterating cursor:\n" + err.Error(),
		})
	}

	// decrypt all tasklists' title and content
	for iTaskList, taskList := range result {
		result[iTaskList].Title, err = utils.AESDecrypt(taskList.Title)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "error while decrypting tasklist title:\n" + err.Error(),
			})
		}

		for iTask, task := range taskList.Tasks {
			result[iTaskList].Tasks[iTask].Content, err = utils.AESDecrypt(task.Content)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "error while decrypting task content:\n" + err.Error(),
				})
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(
		result,
	)
}

func AddTaskList(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	type Input struct {
		Title string `json:"title"`
	}

	var input Input

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
		})
	}

	taskListID, err := utils.CreateTaskList(email, input.Title, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"taskListID": taskListID,
	})
}

func DeleteTaskList(c *fiber.Ctx) error {
	taskLists := config.DB.Collection("task_lists")

	// Convert the param taskListID from string to ObjectID
	taskListID, err := primitive.ObjectIDFromHex(c.Params("taskListID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid tasklist id",
		})
	}

	filter := bson.M{
		"_id":   taskListID,
		"owner": c.Locals("email"),
	}
	result, err := taskLists.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "tasklist not found / permission denied",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
