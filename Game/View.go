package Game

import (
	"fmt"
	"github.com/AlexanderSychev/go_mines/flow"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const cellButtonSize = 24

// ---------------------------------------------------------------------------------------------------------------------
// "View" type definition
// ---------------------------------------------------------------------------------------------------------------------

type View struct {
	root      *gtk.Grid
	timeLabel *gtk.Label
	field     *gtk.Grid
	cells     [][]*CellButton
}

// ---------------------------------------------------------------------------------------------------------------------
// "View" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (v *View) makeCell(x, y int) (*CellButton, error) {
	return NewCellButton(
		func() {
			v.onOpenClick(x, y)
		},
		func() {
			v.onMarkClick(x, y)
		},
	)
}

func (v *View) onOpenClick(x, y int) {
	broker := flow.GetBrokerInstance()
	_, err := broker.SendToActor(flow.ControllerActor, flow.NewMessageWithArgs("OnOpen", x, y))
	if err != nil {
		log.Println(err)
	}
}

func (v *View) onMarkClick(x, y int) {
	broker := flow.GetBrokerInstance()
	_, err := broker.SendToActor(flow.ControllerActor, flow.NewMessageWithArgs("OnMark", x, y))
	if err != nil {
		log.Println(err)
	}
}

func (v *View) SetClock(minutes, seconds int) {
	minutesPrefix := ""
	if minutes < 10 {
		minutesPrefix = "0"
	}
	secondsPrefix := ""
	if seconds < 10 {
		secondsPrefix = "0"
	}
	v.timeLabel.SetLabel(fmt.Sprintf("%s%d:%s%d", minutesPrefix, minutes, secondsPrefix, seconds))
}

func (v *View) SetOpened(x, y, minesAround int) {
	err := v.cells[x][y].SetOpened(minesAround)
	if err != nil {
		log.Println(err)
		return
	}
	v.cells[x][y].Deactivate()
}

func (v *View) SetWithMine(x, y int) {
	err := v.cells[x][y].SetWithMine()
	if err != nil {
		log.Println(err)
		return
	}
	v.cells[x][y].Deactivate()
}

func (v *View) SetMarked(x, y int) {
	err := v.cells[x][y].SetMarked()
	if err != nil {
		log.Println(err)
	}
}

func (v *View) SetClosed(x, y int) {
	err := v.cells[x][y].SetClosed()
	if err != nil {
		log.Println(err)
	}
}

func (v *View) GameOver() {
	for x := 0; x < len(v.cells); x++ {
		for y := 0; y < len(v.cells[x]); y++ {
			v.cells[x][y].Deactivate()
		}
	}
}

func (v *View) Activate() error {
	var err error = nil

	for x := 0; x < len(v.cells); x++ {
		for y := 0; y < len(v.cells[x]); y++ {
			v.cells[x][y], err = v.makeCell(x, y)
			if err != nil {
				return err
			}

			err = v.cells[x][y].Activate()
			if err != nil {
				return err
			}

			v.field.Attach(v.cells[x][y].GetRoot(), x*cellButtonSize, y*cellButtonSize, cellButtonSize, cellButtonSize)
		}
	}

	return err
}

func (v *View) Deactivate() {
	for x := 0; x < len(v.cells); x++ {
		for y := 0; y < len(v.cells[x]); y++ {
			v.cells[x][y].Deactivate()
		}
	}
}

func (v *View) GetWidgets() []gtk.IWidget {
	result := make([]gtk.IWidget, 1, 1)
	result[0] = v.root
	return result
}

func (v *View) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] [go_mines/Game.View]() Does not understand: %v",
		time.Now().Format(time.UnixDate),
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "View" type constructors
// ---------------------------------------------------------------------------------------------------------------------

func NewView(width, height int) (*View, error) {
	root, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	timeLabel, err := gtk.LabelNew("-:-")
	if err != nil {
		return nil, err
	}

	field, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	cells := make([][]*CellButton, height, height)
	for x := 0; x < height; x++ {
		cells[x] = make([]*CellButton, width, width)
	}

	root.Attach(timeLabel, 0, 0, cellButtonSize*width, cellButtonSize)
	root.Attach(field, 0, cellButtonSize, cellButtonSize*width, cellButtonSize*height)

	return &View{
		root:      root,
		timeLabel: timeLabel,
		field:     field,
		cells:     cells,
	}, nil
}
