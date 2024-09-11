package controllers

import (
	"golang-backend/config"
	"golang-backend/models"
	"golang-backend/utils"
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

	// parsing req body to user
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	err = user.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// hashing and storing the password
	passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(passwordBytes)

	_, err = users.InsertOne(c.Context(), user)
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

	// creating the first TaskList for the user
	err = utils.CreateTaskList(user.Email, user.FirstName+"'s Tasks", c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Error while creating TaskList",
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

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Data",
		})
	}

	user := new(models.User)

	filter := bson.M{"_id": data.Email}
	opts := options.FindOne().SetProjection(bson.M{
		"password":   1,
		"first_name": 1,
	})
	err = users.FindOne(c.Context(), filter, opts).Decode(&user)
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

	// The token now has sub: email, and exp: 15mins from login time
	token, err := claims.SignedString([]byte(os.Getenv("LOGIN_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to generate token",
		})
	}

	// Set an HTTPOnly cookie with the token
	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    token,
		HTTPOnly: true,     // Prevents JS access to the token
		Secure:   true,     // Only sent over HTTPS
		SameSite: "Strict", // Mitigates CSRF
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":    true,
		"message":    "Logged in successfully",
		"first_name": user.FirstName,
	})
}
