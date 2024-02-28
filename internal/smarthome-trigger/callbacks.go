package smarthome_trigger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type SceneBodyRequest struct {
	Service     string    `json:"service"`
	ServiceData SceneData `json:"service_data"`
}

type SceneData struct {
	EntityID string `json:"entity_id"`
}

func (s *Service) homeAssistantCommandParser(message string) (*string, error) {
	// parse command
	command, content, err := ParseCommands(message)
	if err != nil {
		// Only early return, since most of the messages are no commands.
		return nil, nil
	}
	if command == LightCommand {
		request := SceneBodyRequest{
			Service: "scene.turn_on",
			ServiceData: SceneData{
				EntityID: "scene." + content,
			},
		}
		byteRequest, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest("POST",
			s.homeAssistantURL+"/api/services/scene", bytes.NewBuffer(byteRequest))
		if err != nil {
			return nil, err
		}
		req.Header.Add("content-type", "application/json")
		req.Header.Add("Authorization", "Bearer "+s.homeAssistantToken)
		response, err := s.apiClient.Do(req)
		if err != nil {
			return nil, err
		}
		fmt.Println("test4: " + response.Status)
		defer response.Body.Close()
		if response.Status != "200" {
			return nil, errors.New(fmt.Sprintf("HA Status Code: %s", response.Status))
		}
	}
	fmt.Println("test5")
	return nil, nil
}
