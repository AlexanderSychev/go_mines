package flow

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	activationSignal      = "activate"
	styleProviderPriority = 400
)

// ---------------------------------------------------------------------------------------------------------------------
// "Application" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Application struct {
	currentPage *Page
	title       string
	cssPath     string
	gtkApp      *gtk.Application
	gtkWindow   *gtk.ApplicationWindow
}

// ---------------------------------------------------------------------------------------------------------------------
// "Application" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (app *Application) dropPage() {
	broker := GetBrokerInstance()

	if app.currentPage != nil {
		app.currentPage.View.Deactivate()
		broker.RemoveActor(ControllerActor)
		broker.RemoveActor(ModelActor)
		broker.RemoveActor(ViewActor)

		for _, widget := range app.currentPage.View.GetWidgets() {
			app.gtkWindow.Remove(widget)
		}
		app.currentPage = nil
	}
}

func (app *Application) reflowPage(page *Page) {
	app.currentPage = page
	broker := GetBrokerInstance()

	if app.currentPage != nil {
		broker.AddActor(ControllerActor, app.currentPage.Controller)
		broker.AddActor(ModelActor, app.currentPage.Model)
		broker.AddActor(ViewActor, app.currentPage.View)

		for _, widget := range app.currentPage.View.GetWidgets() {
			app.gtkWindow.Add(widget)
		}

		err := app.currentPage.View.Activate()

		if err != nil {
			log.Fatal(err)
		}

		app.gtkWindow.ShowAll()
	}
}

// Router message "onRoute" handler
func (app *Application) OnRoute(page Page) {
	app.dropPage()
	app.reflowPage(&page)
}

// Load stylesheet from settled CSS file
func (app *Application) loadStylesheet() error {
	cssProvider, err := gtk.CssProviderNew()

	if err != nil {
		return err
	}

	err = cssProvider.LoadFromPath(app.cssPath)
	if err != nil {
		return err
	}

	display, err := gdk.DisplayGetDefault()
	if err != nil {
		return err
	}

	screen, err := display.GetDefaultScreen()
	if err != nil {
		return err
	}
	gtk.AddProviderForScreen(screen, cssProvider, styleProviderPriority)

	return nil
}

func (app *Application) onActivate(callback func()) {
	var err error

	app.gtkWindow, err = gtk.ApplicationWindowNew(app.gtkApp)
	if err != nil {
		log.Fatal(err)
	}
	app.gtkWindow.SetTitle(app.title)

	if len(app.cssPath) > 0 {
		err = app.loadStylesheet()
		if err != nil {
			log.Printf("Stylesheet loading error: %v", err)
		}
	}

	app.gtkWindow.ShowAll()

	callback()
}

func (app *Application) HandleUnknownMessage(message Message) ([]reflect.Value, error) {
	log.Printf(
		"[go_mines/flow.Application] [%s] Does not understand: %v",
		time.Now().Format(time.UnixDate),
		message,
	)
	return nil, nil
}

func (app *Application) Run(callback func()) {
	app.gtkApp.Connect(activationSignal, func() {
		app.onActivate(callback)
	})
	os.Exit(app.gtkApp.Run(os.Args))
}

// ---------------------------------------------------------------------------------------------------------------------
// "Application" type construction functions
// ---------------------------------------------------------------------------------------------------------------------

func NewApplication(appId string, title string, cssPath string) (*Application, error) {
	gtkApp, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		return nil, err
	}

	result := &Application{
		gtkApp:  gtkApp,
		title:   title,
		cssPath: cssPath,
	}
	GetBrokerInstance().AddActor(ApplicationActor, result)

	return result, nil
}
