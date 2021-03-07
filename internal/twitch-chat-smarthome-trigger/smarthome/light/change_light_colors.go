package light

import (
	"errors"
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/pkg/mqtt"
	"github.com/4ND3R50N/twitch-chat-smarthome-trigger/pkg/project"
	"io/ioutil"
)

func ChangeColor(color Color) (string, error) {

	var payload []byte
	var err error

	switch color {
	case Blue:
		payload, err = ioutil.ReadFile(project.Path + "/assets/smarthome/light/bright_blue.json")
		if err != nil {
			return "", errors.New("unable to load file for color " + Blue.String())
		}
	case Orange:
		payload, err = ioutil.ReadFile(project.Path + "/assets/smarthome/light/bright_orange.json")
		if err != nil {
			return "", errors.New("unable to load file for color " + Blue.String())
		}
	default:
		return "Your color is not implemented yet!", nil
	}
	setLightState(payload)
	return "Successfully changed color", nil
}

func setLightState(payload []byte) {
	// todo: store smarthome devices in seperate file
	// workstation-01
	rs1 := mqtt.BrokerClient.Publish("homeassistant/0x00158d00038ea806/set", 1, false, payload)
	rs1.Wait()
	// workstation-02
	rs2 := mqtt.BrokerClient.Publish("homeassistant/0x00158d00038ebb56/set", 1, false, payload)
	rs2.Wait()
	// workstation-03
	rs3 := mqtt.BrokerClient.Publish("homeassistant/0x00158d00038eca7c/set", 1, false, payload)
	rs3.Wait()
}
