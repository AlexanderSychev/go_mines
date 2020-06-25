package model

import (
	"github.com/AlexanderSychev/go_mines/pubsub"
	"github.com/AlexanderSychev/go_mines/pubsub/ModelEvents"
)

// ---------------------------------------------------------------------------------------------------------------------
// "CustomGame" type definition
// ---------------------------------------------------------------------------------------------------------------------

type CustomGame struct {
	publisher            *pubsub.Publisher
	width, height, mines int
}

// Public methods

func (cg *CustomGame) SetWidth(width int) error {
	err := checkWidth(width)
	if err != nil {
		return err
	}
	cg.width = width
	cg.publisher.Publish(ModelEvents.NewCustomGameEvent(ModelEvents.CustomGameWidthChange, cg.width))
	return nil
}

func (cg *CustomGame) SetHeight(height int) error {
	err := checkHeight(height)
	if err != nil {
		return err
	}
	cg.height = height
	cg.publisher.Publish(ModelEvents.NewCustomGameEvent(ModelEvents.CustomGameHeightChange, cg.height))
	return nil
}

func (cg *CustomGame) SetMines(mines int) error {
	err := checkMines(mines)
	if err != nil {
		return err
	}
	cg.mines = mines
	cg.publisher.Publish(ModelEvents.NewCustomGameEvent(ModelEvents.CustomGameMinesChange, cg.mines))
	return nil
}

// Constructors

func NewCustomGame() *CustomGame {
	return &CustomGame{
		width:  minWidth,
		height: minHeight,
		mines:  minMines,
	}
}
