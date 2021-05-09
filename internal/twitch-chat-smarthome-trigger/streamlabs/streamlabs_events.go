package streamlabs

import (
	"github.com/4ND3R50N/twitch-smarthome-trigger/pkg/streamlabs"
	"log"
)

func PublishStreamLabsCallbacks() {
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := streamlabs.WebsocketConnection.ReadMessage()

			// todo: create use case handler
			if err != nil {
				log.Println("Received message:", err)
				return
			}
			log.Printf("Received message: %s", message)
		}
	}()
}