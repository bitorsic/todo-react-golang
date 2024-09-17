package controllers

import (
	"golang-backend/config"
	"golang-backend/models"
	"golang-backend/utils"
	"time"

	"github.com/gofiber/fiber/v2"
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
			"error": "invalid data",
		})
	}

	err = user.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// hashing and storing the password
	passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(passwordBytes)

	_, err = users.InsertOne(c.Context(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "email is already in use",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while saving to the database:\n" + err.Error(),
		})
	}

	// creating the first TaskList for the user
	_, err = utils.CreateTaskList(user.Email, user.FirstName+"'s Tasks", c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func Login(c *fiber.Ctx) error {
	var users = config.DB.Collection("users")
	data := new(models.User)

	err := c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
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
				"error": "invalid credentials",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while finding in the database:\n" + err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	authToken, err := utils.CreateJWT(user.Email, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while generating auth token\n" + err.Error(),
		})
	}

	refreshToken, err := utils.CreateJWT(user.Email, true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while generating refresh token\n" + err.Error(),
		})
	}

	// Set an HTTPOnly cookie with the refresh token
	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,     // Prevents JS access to the token
		Secure:   true,     // Only sent over HTTPS
		SameSite: "Strict", // Mitigates CSRF
	})

	// Send the auth token in the body, frontend will use it in header
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"first_name": user.FirstName,
		"authToken":  authToken,
	})
}

func TokenRefresh(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing refresh token",
		})
	}

	// firstly verify the refreshToken
	email, err := utils.VerifyJWT(refreshToken, true)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// if refreshToken is valid, generate new authToken for the user
	authToken, err := utils.CreateJWT(email, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while generating auth token\n" + err.Error(),
		})
	}

	// Send the auth token in the body, frontend will use it in header
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"authToken": authToken,
	})
}
