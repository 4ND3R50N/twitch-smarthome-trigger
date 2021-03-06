package smarthome

import (
	"fmt"
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/internal/twitch-chat-smarthome-trigger/smarthome/light"
	"strings"
)

const commandSeparator string = ""
const lightCommand = "change_lights " + commandSeparator

func HandleMessage(message string) string {
	if strings.HasPrefix(message, lightCommand) {
		command := strings.Split(message, commandSeparator)
		fmt.Println("DEBUG: command 1: " + command[1])
		status, err := light.ChangeColor(light.Color(command[1]).ByString(command[1]))

		if err != nil {
			fmt.Printf(err.Error())
		}
		return status
	}
	// add other smarthome commands here
	return ""
}
