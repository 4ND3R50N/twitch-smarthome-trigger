package irc

import (
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/internal/twitch-chat-smarthome-trigger/smarthome"
	"github.com/gempir/go-twitch-irc/v2"
)

func PublishTwitchCallbacks(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		smarthome.HandleMessage(message.Message)
	})

	// add other twitch irc events here
}
