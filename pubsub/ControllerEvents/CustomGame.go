package ControllerEvents

import "github.com/AlexanderSychev/go_mines/pubsub"

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	CustomGameSetWidth  = 100
	CustomGameSetHeight = 101
	CustomGameSetMines  = 102
)

// ---------------------------------------------------------------------------------------------------------------------
// "CustomGame" controller events constructors
// ---------------------------------------------------------------------------------------------------------------------

func NewCustomGameEvent(eventType uint16, payload int) pubsub.Event {
	return pubsub.Event{Type: eventType, Payload: payload}
}
