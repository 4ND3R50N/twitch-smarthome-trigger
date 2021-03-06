package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var BrokerClient mqtt.Client

func Connect(broker string) {
	opts := mqtt.NewClientOptions().AddBroker(broker)

	brokerClient := mqtt.NewClient(opts)
	if token := brokerClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf(token.Error().Error())
	}
}
