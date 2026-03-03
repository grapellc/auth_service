package otp

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Sender interface {
	Send(ctx context.Context, to string, code string) error
}

type ConsoleSender struct{}

func NewConsoleSender() *ConsoleSender {
	return &ConsoleSender{}
}

func (s *ConsoleSender) Send(ctx context.Context, to string, code string) error {
	logrus.WithFields(logrus.Fields{
		"to":   to,
		"code": code,
	}).Warn("Sending OTP to console (sender is gone or down)")
	return nil
}
