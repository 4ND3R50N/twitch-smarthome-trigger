package twitch

import (
	"fmt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/smarthome"
	"github.com/gempir/go-twitch-irc/v2"
)

func PublishTwitchCallbacks(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// todo: add delay for channel messages
		status := smarthome.HandleMessage(message.Message)

		if status != "" {
			// todo: use env var here
			fmt.Println(status)
			client.Say("juel_djteam", status)
		}
	})

	client.OnWhisperMessage(func(message twitch.WhisperMessage) {
		status := smarthome.HandleMessage(message.Message)

		if status != "" {
			// todo: use env var here
			fmt.Println(status)
			client.Say("juel_djteam", status)
		}
	})
}
