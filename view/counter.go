package view

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	counterHeight = 16
	labelWidth    = 32
	spinWidth     = 64
)

// ---------------------------------------------------------------------------------------------------------------------
// "Counter" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Counter struct {
	value          int
	root           *gtk.Grid
	label          *gtk.Label
	spin           *gtk.SpinButton
	onChangeSignal glib.SignalHandle
}

func (counter *Counter) Activate() error {
	onChangeSignal, err := counter.spin.Connect("value-changed", func(_ *gtk.SpinButton, _ interface{}) {
		counter.value = int(counter.spin.GetValue())
	})
	if err != nil {
		return err
	}
	counter.onChangeSignal = onChangeSignal
	return nil
}

func (counter *Counter) Deactivate() {
	counter.spin.HandlerDisconnect(counter.onChangeSignal)
}

func (counter *Counter) GetRootWidget() gtk.IWidget {
	return counter.root
}

func (counter *Counter) GetValue() int {
	return counter.value
}

// Constructor

func NewCounter(min, max int, strLabel string) (*Counter, error) {
	root, err := gtk.GridNew()
	if err != nil {
		return nil, err
	}

	label, err := gtk.LabelNew(strLabel)
	if err != nil {
		return nil, err
	}

	spin, err := gtk.SpinButtonNewWithRange(float64(min), float64(max), 1.0)
	if err != nil {
		return nil, err
	}

	root.Attach(label, 0, 0, labelWidth, counterHeight)
	root.Attach(spin, labelWidth, 0, spinWidth, counterHeight)

	result := &Counter{
		value: min,
		root:  root,
		label: label,
		spin:  spin,
	}

	return result, nil
}
