package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var BrokerClient mqtt.Client

func Connect(broker string, userName string, password string, clientId string) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetUsername(userName).SetPassword(password).SetClientID(clientId)

	BrokerClient := mqtt.NewClient(opts)
	if token := BrokerClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error().Error())
	}
}
