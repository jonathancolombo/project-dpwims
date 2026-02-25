package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client is a wrapper around the MQTT client to manage connections and publish messages.
type Client struct {
	client mqtt.Client
}

const brokerURL = "tcp://localhost:1884"
const clientID = "golang-publisher"
const keepAlive = 3 * time.Second
const connectRetryInterval = 3 * time.Second
const quiesce = 250

// NewClient initializes and returns a new MQTT client with the specified options.
func NewClient() *Client {
	options := mqtt.NewClientOptions()
	options.AddBroker(brokerURL)
	options.SetClientID(clientID)
	options.SetCleanSession(true)
	options.SetAutoReconnect(true)
	options.SetConnectRetry(true)
	options.SetConnectRetryInterval(keepAlive)
	options.SetKeepAlive(connectRetryInterval)

	options.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Println("Connection lost:", err.Error())
	}

	options.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
	}

	client := mqtt.NewClient(options)
	return &Client{client: client}
}

// Connect attempts to connect to the MQTT broker and logs the result.
func (client *Client) Connect() error {
	if token := client.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %s", token.Error())
		return token.Error()
	}
	log.Println("Connected to MQTT broker")
	return nil
}

// Disconnect gracefully disconnects from the MQTT broker and logs the action.
func (client *Client) Disconnect() {
	client.client.Disconnect(quiesce)
	log.Println("Disconnected from MQTT broker")
}

// Subscribe subscribes to a topic with the specified QoS and message handler, and logs the result.
func (client *Client) Subscribe(topic string, qos byte, handler mqtt.MessageHandler) bool {
	token := client.client.Subscribe(topic, qos, handler)
	if token.Wait() && token.Error() != nil {
		log.Printf("Failed to subscribe to topic %s: %s", topic, token.Error())
		return false
	}
	log.Printf("Subscribed to topic %s with QoS %d", topic, qos)
	return true
}

// Publish publishes a message to a topic with the specified QoS and retained flag, and logs the result.
func (client *Client) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := client.client.Publish(topic, qos, retained, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Failed to publish message: %s", token.Error())
		return token.Error()
	}
	log.Printf("Message published to topic %s: %v", topic, payload)
	return nil
}
