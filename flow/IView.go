package flow

import (
	"github.com/gotk3/gotk3/gtk"
	"reflect"
)

// ---------------------------------------------------------------------------------------------------------------------
// "IView" interface definition
// ---------------------------------------------------------------------------------------------------------------------

// Common interface for top-level application views (which should be render as root widgets)
type IView interface {
	// Inherited from "IActor".
	// Special method to handle unknown message. This can be useful for logging or something third party actions.
	// Please, do not use it as "universal" message handler.
	// Only FATAL error should be returned (application will fail with panic), otherwise should return "nil"
	HandleUnknownMessage(message Message) ([]reflect.Value, error)
	// Returns all widgets to attach. If there the only one widget to attach - return slice of widget with one item.
	GetWidgets() []gtk.IWidget
	// This method will be called before view rendering - here can be connected GTK signals handlers.
	// Only FATAL error should be returned (application will fail with panic), otherwise should return "nil"
	Activate() error
	// This method will be called after view rendering - here should be disconnected GTK signals handlers
	Deactivate()
}
