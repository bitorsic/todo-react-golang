package utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"task-inator3000/config"

	"golang.org/x/crypto/bcrypt"
)

func GenerateOTP() string {
	result := make([]byte, OTPLength)
	charsetLength := big.NewInt(int64(len(otpCharSet)))

	for i := range result {
		// generate a secure random number in the range of the charset length
		num, _ := rand.Int(rand.Reader, charsetLength)
		result[i] = otpCharSet[num.Int64()]
	}

	return string(result)
}

func AddOTPtoRedis(otp string, email string, c context.Context) error {
	key := otpKeyPrefix + email

	// hashing the OTP
	data, _ := bcrypt.GenerateFromPassword([]byte(otp), 10)

	// storing otp with expiry
	err := config.RedisClient.Set(c, key, data, otpExp).Err()
	if err != nil {
		return err
	}

	return nil
}

func SendOTP(otp string, recipient string) error {
	sender := os.Getenv("SMTP_EMAIL")
	client := config.SMTPClient

	// setting the sender
	err := client.Mail(sender)
	if err != nil {
		return err
	}

	// set recipient
	err = client.Rcpt(recipient)
	if err != nil {
		return err
	}

	// start writing email
	writeCloser, err := client.Data()
	if err != nil {
		return err
	}

	// contents of the email
	msg := fmt.Sprintf(emailTemplate, recipient, otp)

	// write the email
	_, err = writeCloser.Write([]byte(msg))
	if err != nil {
		return err
	}

	// close and send email
	err = writeCloser.Close()
	if err != nil {
		return err
	}

	return nil
}
