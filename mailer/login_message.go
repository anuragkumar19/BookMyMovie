package mailer

import (
	"bytes"
	"context"
	"time"

	"errors"

	"bookmymovie.app/bookmymovie/mailer/templates"
)

type LoginMessageParams struct {
	OTP        string
	Link       string
	IsNew      bool
	Email      string
	ExpiryTime time.Duration
}

func NewLoginMessage(ctx context.Context, params *LoginMessageParams) (Message, error) {
	b := bytes.NewBuffer(nil)

	if err := templates.LoginEmail(&templates.LoginEmailParams{
		IsNew:    params.IsNew,
		OTP:      params.OTP,
		Link:     params.Link,
		ValidFor: params.ExpiryTime.String(),
	}).Render(ctx, b); err != nil {
		return Message{}, errors.Join(errors.New("failed to render login email template"), err)
	}

	body := b.String()
	return Message{
		Info:    "Login otp and link",
		To:      params.Email,
		Subject: "OTP for login to BookMyMovie",
		Body:    body,
		Headers: nil,
	}, nil
}
