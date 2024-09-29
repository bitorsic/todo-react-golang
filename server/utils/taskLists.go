package utils

import (
	"context"
	"errors"
	"task-inator3000/config"
	"task-inator3000/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// takes in email, title of tasklist, and the context
// putting it in utils because the register endpoint creates the first tasklist for user
func CreateTaskList(email string, title string, c context.Context) (string, error) {
	var users = config.DB.Collection("users")
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

	// appending the id of inserted tasklist to the task_lists field of the user
	options := bson.M{
		"$push": bson.M{
			"task_lists": taskList.ID,
		},
	}
	_, err = users.UpdateByID(c, email, options)
	if err != nil {
		err = errors.New("error while updating in the database:\n" + err.Error())
		return "", err
	}

	return taskListID, nil
}
