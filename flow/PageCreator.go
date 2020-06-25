package flow

// ---------------------------------------------------------------------------------------------------------------------
// "Page" and "PageCreator" types definition
// ---------------------------------------------------------------------------------------------------------------------

// Triple of model, view and controller to use on concrete application's page
type Page struct {
	// Page's model
	Model IActor
	/// Page's view
	View IView
	// Page's controller
	Controller IActor
}

// Factory function which creates "Page" on concrete route
type PageCreator func(params interface{}) (Page, error)
