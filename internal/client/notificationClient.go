package client

import (
	"fmt"
	"pro-link-api/internal/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

type NotificationClient struct {
	Config     *config.Config
	SMTPClient *mail.SMTPClient
}

type INotificationClient interface {
	SendVerifyAccountEmail(email string, name string) error
}

func (n *NotificationClient) SendVerifyAccountEmail(email string, name string) error {
	verifyAccount := mail.NewMSG()
	verifyAccount.SetFrom(n.Config.Email.From).
		AddTo(email).
		SetSubject("Subject: "+name).
		SetBody(mail.TextHTML, htmlBody)

	fmt.Println("Send Verify Account Email ......... to: " + email)
	err := verifyAccount.Send(n.SMTPClient)

	if err != nil {
		fmt.Println("Error Send Verify Account Email")
		return err
	}

	fmt.Println("Send Verify Account Email Successfully")
	return nil
}

func NewNotificationClient(Config *config.Config, smtpClient *mail.SMTPClient) *NotificationClient {
	return &NotificationClient{
		Config:     Config,
		SMTPClient: smtpClient,
	}
}
