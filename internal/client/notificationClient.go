package client

import (
	"bytes"
	"fmt"
	"html/template"
	"pro-link-api/internal/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

type NotificationClient struct {
	Config     *config.Config
	SMTPClient *mail.SMTPClient
}

type INotificationClient interface {
	SendVerifyAccountEmail(email string, name string) error
}

type VerifyEamilRequest struct {
	Email      string
	Name       string
	VerifyCode string
	Url        string
}

func (n *NotificationClient) SendVerifyAccountEmail(req *VerifyEamilRequest) error {

	template, err := ParseTemplateDir(n.Config.Email.Template.ConfirmAccount, req)
	if err != nil {
		return err
	}

	verifyAccount := mail.NewMSG()
	verifyAccount.SetFrom(n.Config.Email.From).
		AddTo(req.Email).
		SetSubject("Subject: "+req.Name).
		SetBody(mail.TextHTML, template)

	fmt.Println("Send Verify Account Email ......... to: " + req.Email)
	err = verifyAccount.Send(n.SMTPClient)

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

func ParseTemplateDir(dir string, data *VerifyEamilRequest) (string, error) {
	t, err := template.ParseFiles(dir)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil

}
