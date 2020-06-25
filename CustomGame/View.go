package CustomGame

import (
	"fmt"
	"github.com/AlexanderSychev/go_mines/flow"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"reflect"
	"strconv"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	lineHeight = 17
	labelWidth = 128
	spinWidth  = 64
)

// ---------------------------------------------------------------------------------------------------------------------
// "View" type definition
// ---------------------------------------------------------------------------------------------------------------------

type View struct {
	root               *gtk.Grid
	width              *gtk.SpinButton
	widthSignalHandle  glib.SignalHandle
	height             *gtk.SpinButton
	heightSignalHandle glib.SignalHandle
	mines              *gtk.SpinButton
	minesSignalHandle  glib.SignalHandle
	submit             *gtk.Button
	submitSignalHandle glib.SignalHandle
}

// ---------------------------------------------------------------------------------------------------------------------
// "View" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (v *View) SetMaxMines(maxMines int) {
	v.mines.SetRange(float64(minMines), float64(maxMines))
}

func (v *View) String() string {
	const nilVal = "<nil>"
	var width, height, mines = nilVal, nilVal, nilVal

	if v.width != nil {
		width = strconv.FormatInt(int64(v.width.GetValue()), 10)
	}

	if v.height != nil {
		height = strconv.FormatInt(int64(v.height.GetValue()), 10)
	}
	if v.mines != nil {
		mines = strconv.FormatInt(int64(v.mines.GetValue()), 10)
	}
	return fmt.Sprintf("[go_mines/CustomGame.View](width=%s, height=%s, mines=%s)", width, height, mines)
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

func (v *View) GetWidgets() []gtk.IWidget {
	result := make([]gtk.IWidget, 1, 1)
	result[0] = v.root
	return result
}

func (v *View) Activate() error {
	const signal = "value-changed"
	var err error

	v.widthSignalHandle, err = v.width.Connect(signal, func(_ *gtk.SpinButton) {
		message := flow.NewMessageWithArgs("OnWidthChange", int(v.width.GetValue()))
		_, err := flow.GetBrokerInstance().SendToActor(flow.ControllerActor, message)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}

	v.heightSignalHandle, err = v.height.Connect(signal, func(_ *gtk.SpinButton) {
		message := flow.NewMessageWithArgs("OnHeightChange", int(v.height.GetValue()))
		_, err := flow.GetBrokerInstance().SendToActor(flow.ControllerActor, message)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}

	v.minesSignalHandle, err = v.mines.Connect(signal, func(_ *gtk.SpinButton) {
		message := flow.NewMessageWithArgs("OnMinesChange", int(v.mines.GetValue()))
		_, err := flow.GetBrokerInstance().SendToActor(flow.ControllerActor, message)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}

	v.submitSignalHandle, err = v.submit.Connect("clicked", func(_ *gtk.Button) {
		_, err := flow.GetBrokerInstance().SendToActor(flow.ControllerActor, flow.NewMessage("OnSubmit"))
		if err != nil {
			log.Println(err)
		}
	})
	return nil
}

func (v *View) Deactivate() {
	v.width.HandlerDisconnect(v.widthSignalHandle)
	v.height.HandlerDisconnect(v.heightSignalHandle)
	v.mines.HandlerDisconnect(v.minesSignalHandle)
}

// ---------------------------------------------------------------------------------------------------------------------
// "View" type construction function
// ---------------------------------------------------------------------------------------------------------------------

func NewView() (*View, error) {
	root, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	width, err := gtk.SpinButtonNewWithRange(float64(minWidth), float64(maxWidth), 1.0)
	if err != nil {
		return nil, err
	}

	height, err := gtk.SpinButtonNewWithRange(float64(minHeight), float64(maxHeight), 1.0)
	if err != nil {
		return nil, err
	}

	mines, err := gtk.SpinButtonNewWithRange(float64(minMines), float64((minWidth*minHeight)/2), 1.0)
	if err != nil {
		return nil, err
	}

	widthLabel, err := gtk.LabelNew("Width:")
	if err != nil {
		return nil, err
	}

	heightLabel, err := gtk.LabelNew("Height:")
	if err != nil {
		return nil, err
	}

	minesLabel, err := gtk.LabelNew("Mines:")
	if err != nil {
		return nil, err
	}

	submit, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}
	submit.SetLabel("Submit")

	root.SetColumnSpacing(1)
	root.Attach(widthLabel, 0, 0, labelWidth, lineHeight)
	root.Attach(width, labelWidth, 0, spinWidth, lineHeight)
	root.Attach(heightLabel, 0, lineHeight, labelWidth, lineHeight)
	root.Attach(height, labelWidth, lineHeight, spinWidth, lineHeight)
	root.Attach(minesLabel, 0, 2*lineHeight, labelWidth, lineHeight)
	root.Attach(mines, labelWidth, 2*lineHeight, spinWidth, lineHeight)
	root.Attach(submit, 0, 3*lineHeight, labelWidth+spinWidth, lineHeight)

	view := &View{
		root:   root,
		width:  width,
		height: height,
		mines:  mines,
		submit: submit,
	}

	return view, nil
}
