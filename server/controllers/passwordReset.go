package controllers

import (
	"task-inator3000/config"
	"task-inator3000/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendPasswordResetEmail(c *fiber.Ctx) error {
	type Input struct {
		Email string `json:"email"`
	}

	var input Input

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
		})
	}

	// check if user with email exists
	users := config.DB.Collection("users")
	filter := bson.M{"_id": input.Email}
	err = users.FindOne(c.Context(), filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user with given email not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error while finding in the database:\n" + err.Error(),
		})
	}

	// generate otp and add it to redis
	otp := utils.GenerateOTP()
	err = utils.AddOTPtoRedis(otp, input.Email, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// send the otp to user through email
	err = utils.SendOTP(otp, input.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func ResetPassword(c *fiber.Ctx) error {
	type Input struct {
		OTP         string `json:"otp"`
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	var input Input

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
		})
	}

	// no input field should be empty
	if input.OTP == "" || input.Email == "" || input.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid data",
		})
	}

	// check redis for otp
	err, isInternalErr := utils.VerifyOTP(input.OTP, input.Email, c.Context())
	if err != nil {
		var code int
		if isInternalErr {
			code = fiber.StatusInternalServerError
		} else {
			code = fiber.StatusUnauthorized
		}

		return c.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = utils.UpdatePassword(input.Email, input.NewPassword, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
