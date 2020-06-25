package view

import (
	"github.com/AlexanderSychev/go_mines/pubsub"

	"github.com/gotk3/gotk3/gtk"
)

// ---------------------------------------------------------------------------------------------------------------------
// "Field" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Field struct {
	publisher *pubsub.Publisher
	grid      *gtk.Grid
}

// Public methods

func (field *Field) Render(win *gtk.ApplicationWindow) {
	if win != nil && field.grid != nil {
		win.Add(field.grid)
	}
}

// Constructor

func NewField(width, height int, publisher *pubsub.Publisher) (*Field, error) {
	grid, gridErr := gtk.GridNew()
	if gridErr != nil {
		return nil, gridErr
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			cell, cellError := NewCell(x, y, publisher)
			if cellError != nil {
				return nil, cellError
			}
			cell.AttachToGrid(grid)
		}
	}

	return &Field{publisher, grid}, nil
}
