package main

import (
	"fmt"

	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/twitch/callbacks"
	"github.com/gempir/go-twitch-irc/v2"
)

func main() {
	fmt.Println("Twitch Chat Smarthome Trigger (call me Mr Johnson) Starting...")
	client := twitch.NewClient("juel_djteam", "oauth:4mwe37ppcvszn69fa5zpu5m6cb26cu")

	fmt.Println("Define callbacks & join channels...")
	callbacks.PublishTwitchCallbacks(client)
	client.Join("juel_djteam")

	fmt.Println("Successfully connected / started!")
	client.Say("juel_djteam", "Mr. Johnson, the butler of JUEL joined the chat. Type in the following commands to interact with the living room :-)")
	err := client.Connect()
	if err != nil {
		panic(err)
	}

}
