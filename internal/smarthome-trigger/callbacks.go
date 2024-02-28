package smarthome_trigger

import (
	"bytes"
	"encoding/json"
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
	fmt.Println("test")
	command, content, err := ParseCommands(message)
	if err != nil {
		// Only early return, since most of the messages are no commands.
		return nil, nil
	}
	fmt.Println("test")
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
		req.Header.Add("Authorization", s.homeAssistantToken)
		response, err := s.apiClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		if response.Status != "200" {
			return nil, err
		}
	}
	return nil, nil
}
