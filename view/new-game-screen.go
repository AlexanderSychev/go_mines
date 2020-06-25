package view

import (
	"github.com/AlexanderSychev/go_mines/pubsub"
	"github.com/gotk3/gotk3/gtk"
)

// ---------------------------------------------------------------------------------------------------------------------
// "NewGame" type definition
// ---------------------------------------------------------------------------------------------------------------------

type NewGameScreen struct {
	grid    *gtk.Grid
	presets [4]View
}

// Public methods

func (screen *NewGameScreen) Activate() error {
	for index, preset := range screen.presets {
		err := preset.Activate()
		if err != nil {
			return err
		}

		left, top := 0, 0
		if index%2 != 0 {
			left = gamePresetSize
		}
		if index > 1 {
			top = gamePresetSize
		}
		screen.grid.Attach(preset.GetRootWidget(), left, top, gamePresetSize, gamePresetSize)
	}

	return nil
}

func (screen *NewGameScreen) Deactivate() {
	for _, preset := range screen.presets {
		preset.Deactivate()
	}
}

func (screen *NewGameScreen) GetRootWidget() gtk.IWidget {
	return screen.grid
}

// Constructor

func NewNewGameScreen(publisher *pubsub.Publisher) (View, error) {
	grid, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	preset8x8x10, err := NewGamePreset(8, 8, 10, publisher)
	if err != nil {
		return nil, err
	}

	preset16x16x40, err := NewGamePreset(16, 16, 40, publisher)
	if err != nil {
		return nil, err
	}

	preset30x16x99, err := NewGamePreset(30, 16, 99, publisher)
	if err != nil {
		return nil, err
	}

	custom, err := NewCustomGamePreset(publisher)
	if err != nil {
		return nil, err
	}

	presets := [4]View{preset8x8x10, preset16x16x40, preset30x16x99, custom}

	result := &NewGameScreen{
		grid:    grid,
		presets: presets,
	}

	return result, nil
}
