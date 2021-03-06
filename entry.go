package main

import (
	"fmt"
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/mqtt"
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/twitch/irc"
)

func main() {
	fmt.Println("Define callbacks...")
	irc.PublishTwitchCallbacks(irc.Client)
	fmt.Println("Join channels...")
	irc.Client.Join("juel_djteam")

	fmt.Printf("Connect to mqtt...")
	mqtt.Connect("tcp://localhost:1883")

	fmt.Println("Start running...")
	irc.Client.Say("juel_djteam", "Mr. Johnson is ONLINE!")
	err := irc.Client.Connect()
	if err != nil {
		panic(err)
	}
}
