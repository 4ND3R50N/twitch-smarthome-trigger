package streamlabs

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

var WebsocketConnection *websocket.Conn

func CreateWebsocket(accessToken string, websocketAddress *string) {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socketUrl := url.URL{Scheme: "ws", Host: *websocketAddress, Path: "?token=" + accessToken}
	log.Printf("connecting to %s", socketUrl.String())

	WebsocketConnection, _, err := websocket.DefaultDialer.Dial(socketUrl.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer WebsocketConnection.Close()
}
