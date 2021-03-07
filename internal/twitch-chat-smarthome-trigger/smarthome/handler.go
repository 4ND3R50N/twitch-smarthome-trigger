package smarthome

import (
	"fmt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/smarthome/light"
	"strings"
)

const commandSeparator string = " "
const lightCommand = "change_lights" + commandSeparator

func HandleMessage(message string) string {
	if strings.HasPrefix(message, lightCommand) {
		changeLights(message)
	}
	// add other smarthome commands here
	return ""
}

func changeLights(message string) string {
	command := strings.Split(message, commandSeparator)
	status, err := light.ChangeColor(light.Color(command[1]))

	if err != nil {
		fmt.Printf(err.Error())
		return "Ups, something went wrong changing the light colors!"
	}
	return status
}
