package Game

import "github.com/AlexanderSychev/go_mines/flow"

func PageCreator(params interface{}) (flow.Page, error) {
	arrParams := params.([3]int)
	width := arrParams[0]
	height := arrParams[1]
	mines := arrParams[2]
	controller := NewController()
	model := NewModel(width, height, mines)
	view, err := NewView(width, height)
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
