package mqtt

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTClient(brokerURL string) mqtt.Client {
	options := mqtt.NewClientOptions()
	options.AddBroker(brokerURL)
	options.SetClientID("trains-service")
	options.SetAutoReconnect(true)

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection failed: %v", token.Error())
	}
	log.Println("MQTT connection established")
	return client
}
