package config

import (
	"crypto/tls"
	"net/smtp"
	"os"
)

var SMTPClient *smtp.Client

func SMTPConnect() {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	smtpAuth := smtp.PlainAuth("", email, password, host)

	// connect to smtp server
	client, err := smtp.Dial(host + ":" + port)
	if err != nil {
		panic(err)
	}

	SMTPClient = client
	client = nil

	// initiate TLS handshake
	err = SMTPClient.StartTLS(&tls.Config{ServerName: host})
	if err != nil {
		panic(err)
	}

	// authenticate
	err = SMTPClient.Auth(smtpAuth)
	if err != nil {
		panic(err)
	}
}
