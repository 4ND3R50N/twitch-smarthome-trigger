package smarthome_trigger

import (
	"fmt"
	"github.com/4ND3R50N/twitch-smarthome-trigger/internal/twitch"
	"net/http"
	"time"
)

type Service struct {
	homeAssistantURL   string
	homeAssistantToken string
	apiClient          http.Client
	twitchTrigger      *twitch.Trigger
}

func NewService(twitchTrigger *twitch.Trigger, token string, url string) *Service {
	return &Service{
		homeAssistantURL:   url,
		homeAssistantToken: token,
		apiClient:          http.Client{Timeout: time.Duration(1) * time.Second},
		twitchTrigger:      twitchTrigger,
	}
}

func (s *Service) Run() error {
	s.twitchTrigger.RegisterPrivateMessageCallbacks(s.homeAssistantCommandParser)
	fmt.Println("Run Service...")
	if err := s.twitchTrigger.Run(); err != nil {
		return err
	}
	return nil
}
