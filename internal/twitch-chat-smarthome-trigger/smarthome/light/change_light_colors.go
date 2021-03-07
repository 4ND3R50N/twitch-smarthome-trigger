package light

import (
	"errors"
	"fmt"
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/pkg/project"
	"io/ioutil"
	"strconv"

	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/pkg/mqtt"
)

func ChangeColor(color Color) (string, error) {
	if color == Blue {
		// todo: make function for sending payload by color
		// todo: add real payload
		file, err := ioutil.ReadFile(project.Path + "/assets/smarthome/light/bright_blue.json")

		if err != nil {
			return "", errors.New("unable to load file for color " + Blue.String())
		}
		fmt.Println("Until here everything works! Mqtt connected: " + strconv.FormatBool(mqtt.BrokerClient.IsConnected()))
		// todo: adjust topic name
		// todo: store smarthome devices in seperate file
		// workstation-01
		fmt.Println("Send message payload" + string(file))
		rs1 := mqtt.BrokerClient.Publish("homeassist/0x00158d00038ea806/set", 1, false, file)
		rs1.Wait()
		// workstation-02
		fmt.Println("Send message")
		rs2 := mqtt.BrokerClient.Publish("homeassist/0x00158d00038ebb56/set", 1, false, file)
		rs2.Wait()
		// workstation-03
		fmt.Println("Send message")
		rs3 := mqtt.BrokerClient.Publish("homeassist/0x00158d00038eca7c/set", 1, false, file)
		rs3.Wait()
		return "Room color changed to " + Blue.String(), nil
	}

	return "Your color is not implemented yet!", nil
}
