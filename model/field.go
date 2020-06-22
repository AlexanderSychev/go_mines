package model

import (
	"fmt"
	"log"
	"os"

	"github.com/AlexanderSychev/go_mines/pubsub"
)

//------------------------------------------------------------------------------
// Constants
//------------------------------------------------------------------------------

const (
	minWidth  = 8
	minHeight = 8
	minMines  = 10
)

//------------------------------------------------------------------------------
// "Field" type definition
//------------------------------------------------------------------------------

type Field struct {
	logger    *log.Logger
	publisher *pubsub.Publisher
	cells     [][]*Cell
}

// Private methods

func (field *Field) checkCoords(x, y int) error {
	var width, height int = field.Width(), field.Height()
	if x < 0 || x >= height {
		return NewCellOpenError(fmt.Sprintf("x=%d is out of range 0..%d", x, height-1))
	}
	if y < 0 || y >= width {
		return NewCellOpenError(fmt.Sprintf("y=%d is out of range 0..%d", y, width-1))
	}
	return nil
}

func (field *Field) getNeighbors(x, y int) UniqueCoordinatesSet {
	width := field.Width()
	height := field.Height()

	result := NewUniqueCoordinatesSet()

	// Topper line (lesser x)
	if x > 0 {
		if y > 0 {
			result.Add(NewCoordinates(x-1, y-1))
		}
		result.Add(NewCoordinates(x-1, y))
		if y < width-1 {
			result.Add(NewCoordinates(x-1, y+1))
		}
	}

	// Same line (same x)
	if y > 0 {
		result.Add(NewCoordinates(x, y-1))
	}
	if y < width-1 {
		result.Add(NewCoordinates(x, y+1))
	}

	if x < height-1 {
		if y > 0 {
			result.Add(NewCoordinates(x+1, y-1))
		}
		result.Add(NewCoordinates(x+1, y))
		if y < width-1 {
			result.Add(NewCoordinates(x+1, y+1))
		}
	}

	return result
}

func (field *Field) getMinesAround(x, y int) int {
	coords := field.getNeighbors(x, y)
	var result int = 0
	for i := 0; i < coords.Length(); i++ {
		coord := coords.Get(i)
		if field.hasMine(coord.X(), coord.Y()) {
			result++
		}
	}
	return result
}

func (field *Field) getEmptyNotOpenedNeighbors(x, y int) UniqueCoordinatesSet {
	neighbors := field.getNeighbors(x, y)
	result := NewUniqueCoordinatesSet()
	for i := 0; i < neighbors.Length(); i++ {
		coord := neighbors.Get(i)
		hasMine := field.hasMine(coord.X(), coord.Y())
		isOpened := field.isOpened(coord.X(), coord.Y())
		if !hasMine && !isOpened {
			result.Add(coord)
		}
	}
	return result
}

func (field *Field) getCell(x, y int) *Cell {
	return field.cells[x][y]
}

func (field *Field) isMarked(x, y int) bool {
	return field.getCell(x, y).IsMarked()
}

func (field *Field) hasMine(x, y int) bool {
	return field.getCell(x, y).HasMine()
}

func (field *Field) isOpened(x, y int) bool {
	return field.getCell(x, y).IsOpened()
}

func (field *Field) openInternal(x, y int) bool {
	if !field.isOpened(x, y) {
		field.getCell(x, y).Open()
		return true
	}

	return false
}

func (field *Field) toggleMarkInternal(x, y int) bool {
	return field.getCell(x, y).ToggleMark()
}

// Public methods

func (field Field) Height() int {
	return len(field.cells)
}

func (field Field) Width() int {
	return len(field.cells[0])
}

func (field *Field) ToggleMark(x, y int) (bool, error) {
	checkErr := field.checkCoords(x, y)
	if checkErr != nil {
		return false, checkErr
	}

	toggled := field.toggleMarkInternal(x, y)

	if toggled {
		payload := pubsub.CellToggleMarkEventPayload{
			X:        x,
			Y:        y,
			IsMarked: field.isMarked(x, y),
		}
		defer field.publisher.Publish(pubsub.NewCellToggleMarkEvent(payload))
	}

	return toggled, nil
}

func (field *Field) Open(x, y int) (bool, error) {
	err := field.checkCoords(x, y)
	if err != nil {
		return false, err
	}

	wasOpened := field.openInternal(x, y)

	if wasOpened {
		minesAround := field.getMinesAround(x, y)
		events := make([]pubsub.Event, 1)
		events[0] = pubsub.NewCellOpenedEvent(pubsub.CellOpenedEventPayload{
			X:           x,
			Y:           y,
			HasMine:     field.cells[x][y].HasMine(),
			IsMarked:    field.cells[x][y].IsMarked(),
			MinesAround: minesAround,
		})

		if minesAround == 0 {
			neighbors := field.getEmptyNotOpenedNeighbors(x, y)
			for i := 1; i < neighbors.Length(); i++ {
				neighbor := neighbors.Get(i)
				opened := field.openInternal(neighbor.X(), neighbor.Y())
				if opened {
					event := pubsub.NewCellOpenedEvent(pubsub.CellOpenedEventPayload{
						X:           neighbor.X(),
						Y:           neighbor.Y(),
						HasMine:     field.hasMine(neighbor.X(), neighbor.Y()),
						IsMarked:    field.isMarked(neighbor.X(), neighbor.Y()),
						MinesAround: field.getMinesAround(neighbor.X(), neighbor.Y()),
					})
					events = append(events, event)
				}
			}
		}

		defer field.publisher.BulkPublish(events)
	}

	return wasOpened, nil
}

func (field *Field) OpenAll() {
	for x, line := range field.cells {
		for y, _ := range line {
			field.Open(x, y)
		}
	}
}

func (field *Field) HandleEvent(event pubsub.Event) {
	payload, hasPayload := event.Payload.(pubsub.CellButtonClickEventPayload)
	if hasPayload {
		switch event.Type {
		case pubsub.CELL_BUTTON_LEFT_CLICK_EVENT:
			opened, err := field.Open(payload.X, payload.Y)
			if err != nil {
				field.logger.Println(err)
			} else if !opened {
				field.logger.Printf(
					"Cell (%d, %d) wasn't opened for some reason\n",
					payload.X,
					payload.Y,
				)
			}
		case pubsub.CELL_BUTTON_RIGHT_CLICK_EVENT:
			toggled, err := field.ToggleMark(payload.X, payload.Y)
			if err != nil {
				field.logger.Println(err)
			} else if !toggled {
				field.logger.Printf(
					"Cell (%d, %d) wasn't toggled for some reason\n",
					payload.X,
					payload.Y,
				)
			}
		default:
			field.logger.Printf("Unknown event type: %d\n", event.Type)
		}
	} else {
		field.logger.Printf("Wrong event payload: %v\n", payload)
	}
}

// Constructor

func NewField(width, height, mines int, publisher *pubsub.Publisher) (*Field, error) {
	if width < minWidth {
		return nil, NewFieldCreationError(fmt.Sprintf(
			"Field width must be at least %d cells",
			minWidth,
		))
	}

	if height < minHeight {
		return nil, NewFieldCreationError(fmt.Sprintf(
			"Field height must be at least %d cells",
			minHeight,
		))
	}

	if mines < minMines {
		return nil, NewFieldCreationError(fmt.Sprintf(
			"Field must contain at least %d mines",
			minMines,
		))
	}

	minesCoords := RandomUniqueCoordinatesSet(width, height, mines)
	cells := make([][]*Cell, height)

	for i := 0; i < height; i++ {
		cells[i] = make([]*Cell, width)
		for j := 0; j < width; j++ {
			cells[i][j] = NewCell(minesCoords.Contains(NewCoordinates(i, j)))
		}
	}

	logger := log.New(os.Stdout, "[model.Field]", log.LstdFlags)
	field := Field{logger, publisher, cells}
	var subscriber pubsub.Subscriber = &field
	publisher.Subscribe(pubsub.CELL_BUTTON_LEFT_CLICK_EVENT, &subscriber)
	publisher.Subscribe(pubsub.CELL_BUTTON_RIGHT_CLICK_EVENT, &subscriber)

	return &field, nil
}

// Finalizer

func FinalizeField(field *Field) bool {
	var success bool = false
	if field != nil {
		if field.publisher != nil {
			var subscriber pubsub.Subscriber = field
			field.publisher.Unsubscribe(pubsub.CELL_BUTTON_LEFT_CLICK_EVENT, &subscriber)
			field.publisher.Unsubscribe(pubsub.CELL_BUTTON_RIGHT_CLICK_EVENT, &subscriber)
			field.publisher = nil
		}
		field.cells = nil
		field.logger = nil
		success = true
	} else {
		log.Println("Field pointer is <nil>")
	}
	return success
}
