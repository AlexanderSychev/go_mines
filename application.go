package main

import (
	"log"
	"os"

	"github.com/AlexanderSychev/go_mines/model"
	"github.com/AlexanderSychev/go_mines/pubsub"
	"github.com/AlexanderSychev/go_mines/view"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// -----------------------------------------------------------------------------
// Constants
// -----------------------------------------------------------------------------

const (
	APP_ID    = "ru.alexander-sychev.go_mines"
	APP_TITLE = "Go Mines"
)

// -----------------------------------------------------------------------------
// "Application" type definition
// -----------------------------------------------------------------------------

type Application struct {
	gtkApp         *gtk.Application
	gtkAppWindow   *gtk.ApplicationWindow
	gtkCssProvider *gtk.CssProvider
	publisher      *pubsub.Publisher
	fieldModel     *model.Field
	fieldView      *view.Field
}

// Methods

func (app *Application) onActivate() {
	gtkAppWindow, err := gtk.ApplicationWindowNew(app.gtkApp)
	if err != nil {
		log.Fatal(err)
	}
	app.gtkAppWindow = gtkAppWindow
	app.gtkAppWindow.SetTitle(APP_TITLE)
	app.gtkAppWindow.SetDefaultSize(640, 480)

	gtkCssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal(err)
	}
	app.gtkCssProvider = gtkCssProvider

	err = app.gtkCssProvider.LoadFromPath("assets/styles.css")
	if err != nil {
		log.Fatal(err)
	}

	display, err := gdk.DisplayGetDefault()
	if err != nil {
		log.Fatal(err)
	}

	screen, err := display.GetDefaultScreen()
	if err != nil {
		log.Fatal(err)
	}
	gtk.AddProviderForScreen(screen, app.gtkCssProvider, 400)

	fieldModel, fieldModelErr := model.NewField(8, 8, 10, app.publisher)
	if fieldModelErr != nil {
		log.Fatal(fieldModelErr)
	}
	app.fieldModel = fieldModel

	fieldView, fieldViewError := view.NewField(8, 8, app.publisher)
	if fieldViewError != nil {
		log.Fatal(fieldViewError)
	}
	app.fieldView = fieldView

	app.fieldView.Render(app.gtkAppWindow)
	app.gtkAppWindow.ShowAll()
}

func (app *Application) Run() {
	app.gtkApp.Connect("activate", func() {
		app.onActivate()
	})
	os.Exit(app.gtkApp.Run(os.Args))
}

// Constructor

func NewApplication() (*Application, error) {
	publisher := pubsub.NewPublisher()

	gtkApp, gtkAppErr := gtk.ApplicationNew(APP_ID, glib.APPLICATION_FLAGS_NONE)
	if gtkAppErr != nil {
		return nil, gtkAppErr
	}

	application := &Application{gtkApp, nil, nil, publisher, nil, nil}

	return application, nil
}
