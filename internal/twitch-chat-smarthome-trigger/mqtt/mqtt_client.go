package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var BrokerClient mqtt.Client

func Connect(broker string) {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetUsername("naokiii").SetPassword("bringMoflv45").SetClientID("twitch-chat-smarthome-trigger")

	brokerClient := mqtt.NewClient(opts)
	if token := brokerClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error().Error())
	}
}
