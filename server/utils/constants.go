package utils

import "time"

const (
	authTokenExp       = time.Minute * 10
	refreshTokenExp    = time.Hour * 24 * 30 // 1 month
	blacklistKeyPrefix = "blacklisted:"
	otpKeyPrefix       = "password-reset:"
	otpExp             = time.Minute * 10
	otpCharSet         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	emailTemplate      = "To: %s\r\n" +
		"Subject: Task-inator 3000 Password Reset\r\n" +
		"\r\n" +
		"Your OTP for password reset is %s\r\n"

	// public because needed for testing
	OTPLength = 10
)
