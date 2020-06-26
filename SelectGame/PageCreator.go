package SelectGame

import "github.com/AlexanderSychev/go_mines/flow"

func PageCreator(_ interface{}) (flow.Page, error) {
	controller := NewController()
	model := NewModel()
	view, err := NewView()
	if err != nil {
		return flow.Page{}, err
	} else {
		return flow.Page{
			Controller: controller,
			Model:      model,
			View:       view,
		}, nil
	}
}
