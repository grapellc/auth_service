package otp

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
)

type CompositeSender struct {
	smsSender   Sender
	emailSender Sender
}

func NewCompositeSender(smsSender Sender, emailSender Sender) *CompositeSender {
	return &CompositeSender{
		smsSender:   smsSender,
		emailSender: emailSender,
	}
}

func (s *CompositeSender) Send(ctx context.Context, to string, code string) error {
	// Simple heuristic: if it contains '@', it's an email
	if strings.Contains(to, "@") {
		if s.emailSender != nil {
			return s.emailSender.Send(ctx, to, code)
		}
		// If no email sender configured, log error and return
		logrus.WithField("to", to).Error("Email OTP sender is gone or not configured")
		return nil
	}

	// Default to SMS
	if s.smsSender != nil {
		return s.smsSender.Send(ctx, to, code)
	}
	logrus.WithField("to", to).Error("SMS OTP sender is gone or not configured")
	return nil
}
