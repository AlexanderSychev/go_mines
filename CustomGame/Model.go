package CustomGame

import (
	"fmt"
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	widthIndex  = 0
	heightIndex = 1
	minesIndex  = 2
)

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Model [3]int

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (model *Model) GetWidth() int {
	return model[widthIndex]
}

func (model *Model) SetWidth(width int) {
	model[widthIndex] = width
}

func (model *Model) GetHeight() int {
	return model[heightIndex]
}

func (model *Model) SetHeight(height int) {
	model[heightIndex] = height
}

func (model *Model) GetMines() int {
	return model[minesIndex]
}

func (model *Model) SetMines(mines int) {
	model[minesIndex] = mines
}

func (model *Model) String() string {
	return fmt.Sprintf(
		"[go_mines/CustomGame.Model](width=%d, height=%d, mines=%d)",
		model.GetWidth(),
		model.GetHeight(),
		model.GetMines(),
	)
}

func (model *Model) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] %v Does not understand: %v",
		time.Now().Format(time.UnixDate),
		model,
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Model" type construction function
// ---------------------------------------------------------------------------------------------------------------------

func NewModel() *Model {
	return &Model{minWidth, minHeight, minMines}
}
