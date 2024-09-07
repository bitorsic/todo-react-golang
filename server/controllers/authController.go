package controllers

import (
	"context"
	"fmt"
	"golang-backend/config"
	"golang-backend/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		fmt.Print(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(passwordBytes)

	result, err := config.Users.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(409).JSON(fiber.Map{
			"success": false,
			"message": "Email is already in use",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "User created",
		"data":    result,
	})
}
func Login(c *fiber.Ctx) error {
	data := new(models.User)

	if err := c.BodyParser(data); err != nil {
		fmt.Print(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	user := new(models.User)

	filter := bson.D{{"_id", data.Email}}
	err := config.Users.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("LOGIN_KEY")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Logged in successfully",
		"token":   token,
	})
}
