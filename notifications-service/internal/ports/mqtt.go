package ports

type MqttPublisher interface {
	Publish(topic string, qos byte, retained bool, payload interface{}) error
}
