package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskList struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title string             `bson:"title" json:"title"`
	Owner string             `bson:"owner" json:"owner"`
	Tasks []string           `bson:"tasks,omitempty" json:"tasks"`
}

func (tl *TaskList) Validate() error {
	if tl.Title == "" {
		return errors.New("title cannot be empty")
	}

	return nil
}
