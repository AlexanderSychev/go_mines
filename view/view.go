package view

import "github.com/gotk3/gotk3/gtk"

// ---------------------------------------------------------------------------------------------------------------------
// "View" interface definition
// ---------------------------------------------------------------------------------------------------------------------

type View interface {
	// View activation handler - here must be a GTK signals connection, styling modifications and same things
	Activate() error
	// View deactivation handler - here must be a
	Deactivate()
	// Returns root widget - to attach on parent widget, not to work with it directly!
	GetRootWidget() gtk.IWidget
}
