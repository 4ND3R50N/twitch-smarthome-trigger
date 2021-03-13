package smarthome

import (
	"fmt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch-chat-smarthome-trigger/smarthome/light"
	"github.com/4ND3R50N/twitch-smarthome-trigger/pkg/timer"
	"strings"
	"time"
)

const commandSeparator string = " "
const lightCommand = "!change_lights" + commandSeparator

var (
	secondsTimer *timer.SecondsTimer
)

func HandleMessage(message string) string {
	if strings.HasPrefix(message, lightCommand) {
		return changeLights(message)
	}
	// add other smart home commands here
	return ""
}

func changeLights(message string) string {

	remaining := secondsTimer.TimeRemaining()
	if remaining.Seconds() != 0 {
		return "Light change is available after " + fmt.Sprintf("%f", remaining.Seconds())
	}

	command := strings.Split(message, commandSeparator)
	status, err := light.ChangeColor(light.Color(command[1]))

	if err != nil {
		fmt.Printf(err.Error())
		return "Ups, something went wrong changing the light colors!"
	}
	duration, err := time.ParseDuration("10h")

	if err != nil {
		fmt.Printf(err.Error())
		return "Ups, something went wrong changing the light colors!"
	}

	secondsTimer = timer.NewSecondsTimer(duration)
	return status
}
