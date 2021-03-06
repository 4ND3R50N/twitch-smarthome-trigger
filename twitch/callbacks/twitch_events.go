package callbacks

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func PublishTwitchCallbacks(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Print(message.User)
		fmt.Println(message.Message)
		client.Say("juel_djteam", "received message")
	})
}
