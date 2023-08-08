package client

import (
	"fmt"
	"log"
	"pro-link-api/internal/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

type Client struct {
	SMTP *mail.SMTPClient
}

func NewClient(config *config.Config) *Client {
	return &Client{
		SMTP: NewSMTPClient(&config.Email),
	}
}

func NewSMTPClient(config *config.EmailConfig) *mail.SMTPClient {
	fmt.Println("Connect SMTP client...")
	server := mail.NewSMTPClient()
	server.Host = config.Host
	server.Port = config.Port
	server.Username = config.Username
	server.Password = config.Password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println("Error Connecting SMTP client")
		log.Fatal(err)
	}

	return smtpClient
}
