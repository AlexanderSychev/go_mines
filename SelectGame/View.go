package SelectGame

import (
	"fmt"
	"github.com/AlexanderSychev/go_mines/flow"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"reflect"
	"time"
)

const (
	buttonSize        = 128
)

// ---------------------------------------------------------------------------------------------------------------------
// "View" type definition
// ---------------------------------------------------------------------------------------------------------------------

type View struct {
	root                     *gtk.Grid
	games []*gtk.Button
	signals []glib.SignalHandle
}

// ---------------------------------------------------------------------------------------------------------------------
// "View" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (v *View) addGame(index int, game string) error {
	broker := flow.GetBrokerInstance()
	const signal = "clicked"
	const selectMethod = "SelectGame"

	button, err := gtk.ButtonNew()
	if err != nil {
		return err
	}
	button.SetLabel(game)

	handle, err := button.Connect(signal, func(_ *gtk.Button) {
		message := flow.NewMessageWithArgs(selectMethod, index)
		fmt.Println(message)
		_, sendErr := broker.SendToActor(flow.ControllerActor, message)
		if sendErr != nil {
			log.Println(sendErr)
		}
	})
	if err != nil {
		return err
	}

	top, left := 0, 0
	if index%2 != 0 {
		left = buttonSize
	}
	if index > 1 {
		top = buttonSize
	}
	v.root.Attach(button, left, top, buttonSize, buttonSize)

	v.games = append(v.games, button)
	v.signals = append(v.signals, handle)

	return nil
}

func (v *View) String() string {
	return "[go_mines/SelectGame.View]"
}

func (v *View) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] %v Does not understand: %v",
		time.Now().Format(time.UnixDate),
		v,
		message,
	)
	return nil, nil
}

func (v *View) Activate() error {
	const getMethod = "GetGames"
	broker := flow.GetBrokerInstance()
	var err error

	res, err := broker.SendToActor(flow.ControllerActor, flow.NewMessage(getMethod))
	if err != nil {
		return err
	}

	if !res[1].IsNil() {
		return res[1].Interface().(error)
	}

	games := res[0].Interface().([]string)
	for index, game := range games {
		err = v.addGame(index, game)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *View) Deactivate() {
	for index, game := range v.games {
		game.HandlerDisconnect(v.signals[index])
	}
}

func (v *View) GetWidgets() []gtk.IWidget {
	result := make([]gtk.IWidget, 1, 1)
	result[0] = v.root
	return result
}

// ---------------------------------------------------------------------------------------------------------------------
// "View" construction function(s)
// ---------------------------------------------------------------------------------------------------------------------

func NewView() (*View, error) {
	root, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	return &View{
		root:         root,
		games: make([]*gtk.Button, 0),
		signals: make([]glib.SignalHandle, 0),
	}, nil
}
