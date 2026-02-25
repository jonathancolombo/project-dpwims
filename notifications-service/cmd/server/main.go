package main

import (
	"fmt"
)

func main() {

	/*
		options := mqtt.NewClientOptions().AddBroker("tcp://localhost:1884").SetClientID("go_mqtt_client")
		options.SetClientID("golang-publisher")

		connection := mqtt.NewClient(options)
		if token := connection.Connect(); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error().Error())
		}

		for i := 0; i < 5; i++ {
			payload := fmt.Sprintf("Message %d", i)
			token := connection.Publish("test/topic", 0, false, payload)
			token.Wait()
		}
	*/
	fmt.Println("Hello World")
}
