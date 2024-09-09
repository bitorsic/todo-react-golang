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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var users = config.DB.Collection("users")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(passwordBytes)

	_, err := users.InsertOne(context.TODO(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"success": false,
				"message": "Email is already in use",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while saving to the database",
			"err":     err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User Created",
	})
}

func Login(c *fiber.Ctx) error {
	var users = config.DB.Collection("users")
	data := new(models.User)

	if err := c.BodyParser(data); err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	user := new(models.User)

	filter := bson.M{"_id": data.Email}
	opts := options.FindOne().SetProjection(bson.M{"password": 1, "first_name": 1})
	err := users.FindOne(context.TODO(), filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Invalid Credentials",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while finding in the database",
			"err":     err,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":    true,
		"message":    "Logged in successfully",
		"email":      user.Email,
		"first_name": user.FirstName,
		"token":      token,
	})
}
