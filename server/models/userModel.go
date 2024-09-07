package models

type User struct {
	Email    string `bson:"_id" json:"email"`
	Password string `bson:"password" json:"password"`
}
