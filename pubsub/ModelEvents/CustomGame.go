package ModelEvents

import "github.com/AlexanderSychev/go_mines/pubsub"

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	CustomGameWidthChange  = 50
	CustomGameHeightChange = 51
	CustomGameMinesChange  = 52
)

// ---------------------------------------------------------------------------------------------------------------------
// "CustomGame" model change events constructors
// ---------------------------------------------------------------------------------------------------------------------

func NewCustomGameEvent(eventType uint16, payload int) pubsub.Event {
	return pubsub.Event{Type: eventType, Payload: payload}
}
