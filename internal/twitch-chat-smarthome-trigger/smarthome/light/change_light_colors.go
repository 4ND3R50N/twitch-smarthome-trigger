package light

import (
	"errors"
	"io/ioutil"

	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/internal/twitch-chat-smarthome-trigger/mqtt"
)

func ChangeColor(color Color) (string, error) {
	if color == Blue {
		// todo: make function for sending payload by color
		// todo: add real payload
		file, err := ioutil.ReadFile("workstation-light-bulb-blue.json")

		if err != nil {
			return "", errors.New("unable to load file for color " + Blue.String())
		}

		// todo: adjust topic name
		token := mqtt.BrokerClient.Publish("homeassist/FRIENDLY_NAME/set", 1, false, file)
		token.Wait()
		return "Room color changed to " + Blue.String(), nil
	}

	return "Your color is not implemented yet!", nil
}
