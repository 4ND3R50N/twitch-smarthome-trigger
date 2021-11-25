package main

import (
	"fmt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/twitch/irc"
	"github.com/4ND3R50N/twitch-smarthome-trigger/pkg/mqtt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/pkg/project"
)

func main() {
	project.InitializeProjectPath()
	fmt.Println("Define callbacks...")
	irc.PublishTwitchCallbacks(irc.Client)
	fmt.Println("Join channels...")
	// todo: use env vars
	irc.Client.Join("Houseaffair")

	fmt.Println("Connect to mqtt...")
	// todo: use env vars
	mqtt.Connect("tcp://localhost:1883", "naokiii", "bringMoflv45", "twitch-chat-smarthome-trigger")
	fmt.Println("Start running...")
	// todo: use env vars
	irc.Client.Say("Houseaffair", "BOT is ONLINE!")
	err := irc.Client.Connect()
	if err != nil {
		panic(err)
	}
}
