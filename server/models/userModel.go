package models

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Email     string               `bson:"_id" json:"email"`
	FirstName string               `bson:"first_name" json:"first_name"`
	LastName  string               `bson:"last_name,omitempty" json:"last_name"`
	Password  string               `bson:"password" json:"password"`
	TaskLists []primitive.ObjectID `bson:"task_lists,omitempty" json:"task_lists"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)

	if !re.MatchString(u.Email) {
		return errors.New("email invalid")
	}

	if u.FirstName == "" {
		return errors.New("first name cannot be empty")
	}

	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}
