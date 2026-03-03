package messaging

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var (
	MqttClient mqtt.Client
)

func NewMqttClient() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(viper.GetString("emqx.broker"))

	// Append random suffix to avoid ClientId collisions
	clientID := fmt.Sprintf("%s_%d", viper.GetString("emqx.client_id"), time.Now().UnixNano())
	opts.SetClientID(clientID)

	opts.SetUsername(viper.GetString("emqx.username"))
	opts.SetPassword(viper.GetString("emqx.password"))
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetCleanSession(true) // Set to true for ephemeral clients to avoid state issues
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		fmt.Printf("[MQTT] Connected with Client ID: %s\n", clientID)
	})
	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		fmt.Printf("[MQTT] Connection lost: %v\n", err)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	MqttClient = client
	return nil
}

func SubscribeMqtt(topic string, handler mqtt.MessageHandler) error {
	if MqttClient == nil {
		return fmt.Errorf("MQTT client not initialized")
	}
	token := MqttClient.Subscribe(topic, 1, handler)
	token.Wait()
	return token.Error()
}

func CloseMqtt() {
	if MqttClient != nil {
		MqttClient.Disconnect(250)
	}
}
