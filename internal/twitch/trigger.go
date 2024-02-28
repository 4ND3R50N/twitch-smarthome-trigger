package twitch

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
)

type Trigger struct {
	ChannelName string
	client      *twitch.Client
}

type TriggerOpts struct {
	TwitchChannelName string
	TwitchUser        string
	TwitchOAuth       string
}

func NewTrigger(opts TriggerOpts) *Trigger {
	return &Trigger{
		ChannelName: opts.TwitchChannelName,
		client:      NewClient(opts.TwitchUser, opts.TwitchOAuth),
	}
}

func (t *Trigger) RegisterPrivateMessageCallbacks(handleMessage func(message string) (*string, error)) {
	t.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println("Incoming message: " + message.Message)
		feedbackText, err := handleMessage(message.Message)
		if err != nil {
			fmt.Println(fmt.Sprintf("PrivateMessageCallback error: %v", err))
		}
		if feedbackText != nil {
			t.client.Say(t.ChannelName, *feedbackText)
		}
	})
	fmt.Println("PrivateMessageCallbacks registered")
}

func (t *Trigger) RegisterWhisperCallbacks(handleMessage func(message string) (*string, error)) {
	t.client.OnWhisperMessage(func(message twitch.WhisperMessage) {
		fmt.Println("Incoming message: " + message.Message)
		feedbackText, err := handleMessage(message.Message)
		if err != nil {
			fmt.Println(fmt.Sprintf("WhisperCallback error: %v", err))
		}
		if feedbackText != nil {
			t.client.Say(t.ChannelName, *feedbackText)
		}
	})
	fmt.Println("WhisperMessageCallbacks registered")
}

func (t *Trigger) Run() error {
	t.client.Say(t.ChannelName, "SmartHome Trigger is ONLINE!")
	if err := t.client.Connect(); err != nil {
		return err
	}
	return nil
}
