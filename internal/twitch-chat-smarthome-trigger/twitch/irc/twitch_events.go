package irc

import (
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/internal/twitch-chat-smarthome-trigger/smarthome"
	"github.com/gempir/go-twitch-irc/v2"
)

func PublishTwitchCallbacks(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		status := smarthome.HandleMessage(message.Message)

		if status != "" {
			client.Say("juel_djteam", status)
		}
	})

	// add other twitch irc events here
}
