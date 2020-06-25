package view

import (
	"fmt"
	"github.com/AlexanderSychev/go_mines/pubsub"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	gamePresetSize        = 128
	customGamePresetLabel = "Custom game"
)

// ---------------------------------------------------------------------------------------------------------------------
// "GamePreset" type definition
// ---------------------------------------------------------------------------------------------------------------------

type GamePreset struct {
	button             *gtk.Button
	buttonSignalHandle glib.SignalHandle
	onClick            func()
}

// Public methods

func (preset *GamePreset) Activate() error {
	if preset.button != nil {
		buttonSignalHandle, err := preset.button.Connect("button-press-event", preset.onClick)
		if err != nil {
			return err
		}
		preset.buttonSignalHandle = buttonSignalHandle
	}

	return nil
}

func (preset *GamePreset) Deactivate() {
	preset.button.HandlerDisconnect(preset.buttonSignalHandle)
}

func (preset *GamePreset) GetRootWidget() gtk.IWidget {
	return preset.button
}

// Constructors

func NewGamePreset(width, height, mines int, publisher *pubsub.Publisher) (View, error) {
	button, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}

	button.SetLabel(fmt.Sprintf("%dx%d, %d mines", width, height, mines))

	onClick := func() {
		publisher.Publish(pubsub.NewCreateNewFieldEvent(width, height, mines))
	}

	return &GamePreset{button: button, onClick: onClick}, nil
}

func NewCustomGamePreset(publisher *pubsub.Publisher) (View, error) {
	button, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}

	button.SetLabel(customGamePresetLabel)

	onClick := func() {
		publisher.Publish(pubsub.NewCustomGameEvent())
	}

	return &GamePreset{button: button, onClick: onClick}, nil
}
