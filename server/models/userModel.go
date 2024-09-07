package models

type User struct {
	Email     string `bson:"_id" json:"email"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Password  string `bson:"password" json:"password"`
}
