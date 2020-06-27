package Game

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"strconv"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	cellImageSize           = 16
	flagImagePath           = "assets/flag.svg"
	mineImagePath           = "assets/mine.svg"
	cellButtonRootClass     = "go-mines__cell"
	cellButtonWithMineClass = "go-mines__cell_with-mine"
	cellButtonOpenedClass   = "go-mines__cell_opened"
)

// ---------------------------------------------------------------------------------------------------------------------
// "CellButton" type definition
// ---------------------------------------------------------------------------------------------------------------------

type CellButton struct {
	activated    bool
	button       *gtk.Button
	signalHandle glib.SignalHandle
	mineImage    *gtk.Image
	flagImage    *gtk.Image
	onLeftClick  func()
	onRightClick func()
}

// ---------------------------------------------------------------------------------------------------------------------
// "CellButton" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (b *CellButton) SetOpened(minesAround int) error {
	b.flagImage.Hide()
	b.mineImage.Hide()

	label := ""
	if minesAround > 0 {
		label = strconv.FormatInt(int64(minesAround), 10)
	}
	b.button.SetLabel(label)
	b.button.Show()

	styleContext, err := b.button.GetStyleContext()
	if err != nil {
		return err
	}
	styleContext.AddClass(cellButtonOpenedClass)

	return nil
}

func (b *CellButton) SetWithMine() error {
	b.flagImage.Hide()

	b.button.SetImage(b.mineImage)
	b.mineImage.Show()

	styleContext, err := b.button.GetStyleContext()
	if err != nil {
		return err
	}
	styleContext.AddClass(cellButtonWithMineClass)

	return nil
}

func (b *CellButton) SetMarked() error {
	b.mineImage.Hide()

	b.button.SetImage(b.flagImage)
	b.flagImage.Show()

	return nil
}

func (b *CellButton) SetClosed() error {
	b.flagImage.Hide()
	b.mineImage.Hide()

	b.button.SetLabel("")
	return nil
}

func (b *CellButton) Activate() error {
	const signal = "button-press-event"
	const leftButton = 1
	const rightButton = 3
	var err error
	if !b.activated {
		b.signalHandle, err = b.button.Connect(signal, func(_ *gtk.Button, e *gdk.Event) {
			event := gdk.EventButtonNewFromEvent(e)
			if event != nil {
				switch event.Button() {
				case leftButton:
					b.onLeftClick()
				case rightButton:
					b.onRightClick()
				}
			}
		})

		b.activated = true

		return err
	}
	return nil
}

func (b *CellButton) Deactivate() {
	if b.activated {
		b.button.HandlerDisconnect(b.signalHandle)
		b.activated = false
	}
}

func (b *CellButton) GetRoot() gtk.IWidget {
	return b.button
}

// ---------------------------------------------------------------------------------------------------------------------
// "CellButton" type constructors
// ---------------------------------------------------------------------------------------------------------------------

func NewCellButton(onLeftClick func(), onRightClick func()) (*CellButton, error) {
	button, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}

	flagImage, err := gtk.ImageNewFromFile(flagImagePath)
	if err != nil {
		return nil, err
	}
	flagImage.SetSizeRequest(cellImageSize, cellImageSize)

	mineImage, err := gtk.ImageNewFromFile(mineImagePath)
	if err != nil {
		return nil, err
	}
	mineImage.SetSizeRequest(cellImageSize, cellImageSize)

	styleContext, err := button.GetStyleContext()
	if err != nil {
		return nil, err
	}
	styleContext.AddClass(cellButtonRootClass)

	return &CellButton{
		activated:    false,
		button:       button,
		onLeftClick:  onLeftClick,
		onRightClick: onRightClick,
		flagImage:    flagImage,
		mineImage:    mineImage,
	}, nil
}
