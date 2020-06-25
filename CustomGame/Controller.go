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
	minWidth  = 8
	maxWidth  = 100
	minHeight = 8
	maxHeight = 100
	minMines  = 10
)

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Controller struct {
	maxMines int
}

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (c *Controller) recalculateMaxMines() error {
	broker := flow.GetBrokerInstance()

	widthRes, err := broker.SendToActor(flow.ModelActor, flow.NewMessage("GetWidth"))
	if err != nil {
		return err
	}
	width := int(widthRes[0].Int())

	heightRes, err := broker.SendToActor(flow.ModelActor, flow.NewMessage("GetHeight"))
	if err != nil {
		return err
	}
	height := int(heightRes[0].Int())

	c.maxMines = (width * height) / 2

	_, err = broker.SendToActor(flow.ViewActor, flow.NewMessageWithArgs("SetMaxMines", c.maxMines))
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) OnWidthChange(width int) error {
	broker := flow.GetBrokerInstance()

	if width < minWidth {
		width = minWidth
	} else if width > maxWidth {
		width = maxWidth
	}

	_, err := broker.SendToActor(flow.ModelActor, flow.NewMessageWithArgs("SetWidth", width))
	if err != nil {
		return err
	}

	return c.recalculateMaxMines()
}

func (c *Controller) OnHeightChange(height int) error {
	broker := flow.GetBrokerInstance()

	if height < minHeight {
		height = minHeight
	} else if height > maxHeight {
		height = maxHeight
	}

	_, err := broker.SendToActor(flow.ModelActor, flow.NewMessageWithArgs("SetHeight", height))
	if err != nil {
		return err
	}

	return c.recalculateMaxMines()
}

func (c *Controller) OnMinesChange(mines int) error {
	broker := flow.GetBrokerInstance()

	if mines < minMines {
		mines = minMines
	} else if mines > c.maxMines {
		mines = c.maxMines
	}

	_, err := broker.SendToActor(flow.ModelActor, flow.NewMessageWithArgs("SetMines", mines))
	if err != nil {
		return err
	}

	return nil
}

func (c *Controller) OnSubmit() error {
	broker := flow.GetBrokerInstance()

	widthRes, err := broker.SendToActor(flow.ModelActor, flow.NewMessage("GetWidth"))
	if err != nil {
		return err
	}
	width := int(widthRes[0].Int())

	heightRes, err := broker.SendToActor(flow.ModelActor, flow.NewMessage("GetHeight"))
	if err != nil {
		return err
	}
	height := int(heightRes[0].Int())

	minesRes, err := broker.SendToActor(flow.ModelActor, flow.NewMessage("GetMines"))
	if err != nil {
		return err
	}
	mines := int(minesRes[0].Int())

	params := [3]int{width, height, mines}

	_, err = broker.RouteTo("game", params)

	return err
}

func (c *Controller) String() string {
	return fmt.Sprintf("[go_mines/CustomGame.Controller](maxMines=%d)", c.maxMines)
}

func (c *Controller) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] %v Does not understand: %v",
		time.Now().Format(time.UnixDate),
		c,
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type construction function
// ---------------------------------------------------------------------------------------------------------------------

func NewController() *Controller {
	maxMines := (minWidth * minHeight) / 2
	return &Controller{
		maxMines: maxMines,
	}
}
