package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Content string             `bson:"content" json:"content"`
}

func (t *Task) Validate() error {
	if t.Content == "" {
		return errors.New("content cannot be empty")
	}

	return nil
}
