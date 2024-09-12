package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content string             `bson:"content" json:"content"`
}

func (tl *Task) Validate() error {
	if tl.Content == "" {
		return errors.New("content cannot be empty")
	}

	return nil
}
