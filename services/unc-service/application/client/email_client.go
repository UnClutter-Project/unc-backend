package client

import (
	"context"
	"fmt"
	"unc/services/unc-service/config"

	brevo "github.com/getbrevo/brevo-go/lib"
)

type EmailClient interface {
	SendOTP(ctx context.Context, email, username, content string) error
}

type EmailClientImpl struct {
	br *brevo.APIClient
}

func NewEmailClient() EmailClient {
	cfg := brevo.NewConfiguration()
	cfg.AddDefaultHeader("api-key", config.GetConfig().BrevoAPIKey)
	return &EmailClientImpl{
		br: brevo.NewAPIClient(cfg),
	}
}

func (c *EmailClientImpl) SendOTP(ctx context.Context, email, username, message string) error {
	sendSmtpEmail := brevo.SendSmtpEmail{
		Sender: &brevo.SendSmtpEmailSender{
			Name:  "UnClutter",
			Email: "unclutter067@gmail.com",
		},
		To: []brevo.SendSmtpEmailTo{
			{
				Email: email,
				Name:  username,
			},
		},
		Subject:     "Your UnClutter OTP Code",
		HtmlContent: fmt.Sprintf("<h3>%s</h3>", message),
	}

	_, _, err := c.br.TransactionalEmailsApi.SendTransacEmail(ctx, sendSmtpEmail)
	if err != nil {
		return err
	}

	return nil
}
