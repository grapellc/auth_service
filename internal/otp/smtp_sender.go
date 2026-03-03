package otp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"
)

type SMTPSender struct {
	Host      string
	Port      int
	Username  string
	Password  string
	FromEmail string
}

func NewSMTPSender(host string, port int, username, password, fromEmail string) *SMTPSender {
	return &SMTPSender{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		FromEmail: fromEmail,
	}
}

func (s *SMTPSender) Send(ctx context.Context, to string, code string) error {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = s.FromEmail
	headers["To"] = to
	headers["Subject"] = "Баталгаажуулах код" // Verification Code
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + fmt.Sprintf("Таны баталгаажуулах код: %s", code)

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	// Choose connection method based on port
	// Port 465 is implicit SSL/TLS
	if s.Port == 465 {
		return s.sendSSL(addr, to, []byte(message))
	}

	// Port 587 or 25 usually STARTTLS
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	if err := smtp.SendMail(addr, auth, s.FromEmail, []string{to}, []byte(message)); err != nil {
		return fmt.Errorf("failed to send email via SMTP: %w", err)
	}

	logrus.Infof("Sent OTP via SMTP to %s", to)
	return nil
}

func (s *SMTPSender) sendSSL(addr, to string, msg []byte) error {
	// TLS Config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return fmt.Errorf("tls connection failed: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return fmt.Errorf("smtp client creation failed: %w", err)
	}
	defer client.Quit()

	// Auth
	if s.Username != "" && s.Password != "" {
		auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth failed: %w", err)
		}
	}

	// Mail
	if err = client.Mail(s.FromEmail); err != nil {
		return fmt.Errorf("smtp mail command failed: %w", err)
	}

	// Rcpt
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("smtp rcpt command failed: %w", err)
	}

	// Data
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("smtp data command failed: %w", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("writing message body failed: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("closing message body failed: %w", err)
	}

	logrus.Infof("Sent OTP via SMTP (SSL) to %s", to)
	return nil
}
