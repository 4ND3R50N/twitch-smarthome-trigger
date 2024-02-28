package twitch

import (
	"github.com/gempir/go-twitch-irc/v2"
	"sync"
)

var (
	once   sync.Once
	client *twitch.Client
)

func NewClient(user string, oauth string) *twitch.Client {
	once.Do(func() {
		client = twitch.NewClient(user, oauth)
	})
	return client
}
