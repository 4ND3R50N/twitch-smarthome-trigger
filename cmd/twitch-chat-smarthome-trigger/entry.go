package main

import (
	"flag"
	"fmt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/streamlabs"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/twitch"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/twitch/irc"
	"github.com/4ND3R50N/twitch-smarthome-trigger/pkg/mqtt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/pkg/project"
	streamLabsConnection "github.com/4ND3R50N/twitch-smarthome-trigger/pkg/streamlabs"
)

func main() {
	project.InitializeProjectPath()

}

func setupTwitchChatTrigger() {
	fmt.Println("Define callbacks...")
	twitch.PublishTwitchCallbacks(irc.Client)
	fmt.Println("Join channels...")
	// todo: use env vars
	irc.Client.Join("juel_djteam")

	fmt.Println("Connect to mqtt...")
	// todo: use env vars
	mqtt.Connect("tcp://localhost:1883", "naokiii", "bringMoflv45", "twitch-chat-smarthome-trigger")
	fmt.Println("Start running...")
	// todo: use env vars
	irc.Client.Say("juel_djteam", "BOT is ONLINE!")
	err := irc.Client.Connect()
	if err != nil {
		panic(err)
	}
}

func setupStreamLabsTrigger() {
	var accessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbiI6IjM3M0EzMUM2N0FDMjAwODEwNzNBIiwicmVhZF9vbmx5Ijp0cnVlLCJwcmV2ZW50X21hc3RlciI6dHJ1ZSwieW91dHViZV9pZCI6IlVDVVBHbUhkdWFDZkRVUjNqc3ExdE5MdyIsInR3aXRjaF9pZCI6IjUzMjc3MjE3NiJ9.zVuekXxqHj5R-szRTs6w0sptaqfLObnJIxQtPZWCs9s"
	var websocketAddress = flag.String("addr", "https://sockets.streamlabs.com", "streamlabs address")

	streamLabsConnection.CreateWebsocket(accessToken, websocketAddress)
	streamlabs.PublishStreamLabsCallbacks()
}