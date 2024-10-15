package utils

import (
	"context"
	"errors"
	"task-inator3000/config"
	"task-inator3000/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// takes in email, title of tasklist, and the context
// putting it in utils because the register endpoint creates the first tasklist for user
func CreateTaskList(email string, title string, c context.Context) (string, error) {
	var taskLists = config.DB.Collection("task_lists")

	// creating the TaskList for the user
	taskList := models.TaskList{
		Title: title,
		Owner: email,
	}

	err := taskList.Validate()
	if err != nil {
		return "", err
	}

	// encrypt title
	taskList.Title, err = AESEncrypt(taskList.Title)
	if err != nil {
		err = errors.New("error while encrypting title:\n" + err.Error())
		return "", err
	}

	// saving the tasklist to DB
	result, err := taskLists.InsertOne(c, taskList)
	if err != nil {
		err = errors.New("error while saving to the database:\n" + err.Error())
		return "", err
	}

	taskListID := result.InsertedID.(primitive.ObjectID).Hex()

	return taskListID, nil
}
