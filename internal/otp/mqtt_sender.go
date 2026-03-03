package otp

import (
	"context"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type MqttSender struct {
	client mqtt.Client
	topic  string
}

func NewMqttSender(client mqtt.Client, topic string) *MqttSender {
	return &MqttSender{
		client: client,
		topic:  topic,
	}
}

type OtpPayload struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func (s *MqttSender) Send(ctx context.Context, to string, code string) error {
	payload := OtpPayload{
		Phone: to,
		Code:  code,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal OTP payload: %w", err)
	}

	token := s.client.Publish(s.topic, 1, false, data)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to publish OTP message: %w", token.Error())
	}

	logrus.Infof("Sent OTP via MQTT to %s", to)
	return nil
}
