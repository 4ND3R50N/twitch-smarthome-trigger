package smarthome_trigger

import (
	"errors"
	"strings"
)

type Command string

const (
	LightCommand   Command = "!lights"
	UnknownCommand Command = "unknown"
)

// ParseCommands parses and returns the command found in the text. It also returns the rest separately.
// Example: !lights on -> command: "!lights" text: "on"
func ParseCommands(text string) (Command, string, error) {
	content := strings.Split(text, " ")
	if len(content) < 1 {
		return UnknownCommand, "", errors.New("this is no command")
	}
	if strings.HasPrefix(text, string(LightCommand)) {
		return LightCommand, content[1], nil
	}
	// add other smart home commands here
	return UnknownCommand, "", errors.New("unknown command")
}
