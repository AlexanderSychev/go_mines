package view

import (
	"log"
	"os"
	"strconv"

	"github.com/AlexanderSychev/go_mines/pubsub"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// -----------------------------------------------------------------------------
// Constants
// -----------------------------------------------------------------------------

const (
	cellButtonSize = 24
	cellImageSize  = 16
	flagImagePath  = "assets/flag.svg"
	mineImagePath  = "assets/mine.svg"
)

// -----------------------------------------------------------------------------
// "Cell" type definition
// -----------------------------------------------------------------------------

type Cell struct {
	x, y      int
	logger    *log.Logger
	publisher *pubsub.Publisher
	button    *gtk.Button
	flagImage *gtk.Image
	mineImage *gtk.Image
}

// Methods

func (cell *Cell) onClick(event *gdk.EventButton) {
	payload := pubsub.CellButtonClickEventPayload{
		X: cell.x,
		Y: cell.y,
	}
	if event.Button() == 3 {
		cell.publisher.Publish(pubsub.NewCellButtonClickEvent(true, payload))
	} else if event.Button() == 1 {
		cell.publisher.Publish(pubsub.NewCellButtonClickEvent(false, payload))
	}
}

func (cell *Cell) HandleEvent(event pubsub.Event) {
	switch event.Type {
	case pubsub.CELL_TOGGLE_MARK_EVENT:
		payload, payloadOk := event.Payload.(pubsub.CellToggleMarkEventPayload)
		if payloadOk && payload.X == cell.x && payload.Y == cell.y {
			if payload.IsMarked {
				cell.button.SetImage(cell.flagImage)
				cell.flagImage.Show()
			} else {
				cell.flagImage.Hide()
				cell.button.SetLabel("")
			}
		}
	case pubsub.CELL_OPENED_EVENT:
		payload, payloadOk := event.Payload.(pubsub.CellOpenedEventPayload)
		if payloadOk && payload.X == cell.x && payload.Y == cell.y {
			cell.flagImage.Hide()
			if payload.HasMine {
				cell.button.SetImage(cell.mineImage)
				cell.mineImage.Show()
				styleContext, styleContextError := cell.button.GetStyleContext()
				if styleContextError == nil {
					styleContext.AddClass("go-mines__cell_with-mine")
				}
			} else {
				var label string = ""
				if payload.MinesAround > 0 {
					label = strconv.FormatInt(int64(payload.MinesAround), 10)
				}
				cell.button.SetLabel(label)
				styleContext, styleContextError := cell.button.GetStyleContext()
				if styleContextError == nil {
					styleContext.AddClass("go-mines__cell_opened")
				}
			}
		}
	}
}

func (cell *Cell) AttachToGrid(grid *gtk.Grid) {
	grid.Attach(
		cell.button,
		cell.y*cellButtonSize,
		cell.x*cellButtonSize,
		cellButtonSize,
		cellButtonSize,
	)
}

// Constructor

func NewCell(x, y int, publisher *pubsub.Publisher) (*Cell, error) {
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

	button, err := gtk.ButtonNew()
	if err != nil {
		return nil, err
	}

	styleContext, err := button.GetStyleContext()
	if err != nil {
		return nil, err
	}
	styleContext.AddClass("go-mines__cell")

	logger := log.New(os.Stdout, "[view.Cell]", log.LstdFlags)

	view := Cell{x, y, logger, publisher, button, flagImage, mineImage}
	var subscriber pubsub.Subscriber = &view
	publisher.Subscribe(pubsub.CELL_OPENED_EVENT, &subscriber)
	publisher.Subscribe(pubsub.CELL_TOGGLE_MARK_EVENT, &subscriber)

	button.SetLabel("")
	_, connectErr := button.Connect("button-press-event", func(_ *gtk.Button, e *gdk.Event) {
		view.onClick(gdk.EventButtonNewFromEvent(e))
	})

	if connectErr != nil {
		return nil, connectErr
	}

	return &view, nil
}

// Finalizer

func FinalizeCell(cell *Cell) {
}
