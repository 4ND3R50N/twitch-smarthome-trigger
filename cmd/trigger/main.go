package main

import (
	"flag"
	"fmt"
	smarthomeTrigger "github.com/4ND3R50N/twitch-smarthome-trigger/internal/smarthome-trigger"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch"
	"github.com/peterbourgon/ff/v3"
	"os"
)

func main() {
	fs := flag.NewFlagSet("st", flag.ExitOnError)

	triggerOpts := twitch.TriggerOpts{}

	fmt.Println("Collect ENV vars")
	fs.StringVar(&triggerOpts.TwitchChannelName, "twitch-channel-name", "", "Twitch Channel Name")
	fs.StringVar(&triggerOpts.TwitchUser, "twitch-user", "", "Twitch user")
	fs.StringVar(&triggerOpts.TwitchOAuth, "twitch-oauth", "", "Twitch OAuth")
	var (
		homeAssistantToken string
		homeAssistantURL   string
	)
	fs.StringVar(&homeAssistantURL, "home-assistant-url", "", "HomeAssistant URL")
	fs.StringVar(&homeAssistantToken, "home-assistant-api-token", "", "HomeAssistant Token")
	if err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("st")); err != nil {
		panic("unable to parse flags")
	}

	fmt.Println("Initialize services")
	trigger := twitch.NewTrigger(triggerOpts)
	service := smarthomeTrigger.NewService(trigger, homeAssistantToken, homeAssistantURL)

	fmt.Println("Run...")
	if err := service.Run(); err != nil {
		panic(err)
	}
}
