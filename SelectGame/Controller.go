package SelectGame

import (
	"github.com/AlexanderSychev/go_mines/Routing"
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Controller struct{}

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (c *Controller) GetGames() ([]string, error) {
	res, err := flow.GetBrokerInstance().SendToActor(flow.ModelActor, flow.NewMessage("GetAllGames"))
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for i := 0; i < res[0].Len(); i++ {
		game := res[0].Index(i).Interface().(Game)
		result = append(result, game.Label())
	}
	return result, nil
}

func (c *Controller) SelectGame(index int) {
	res, _ := flow.GetBrokerInstance().SendToActor(flow.ModelActor, flow.NewMessageWithArgs("GetGame", index))
	game := res[0].Interface().(Game)
	if !game.IsCustom() {
		params := [3]int{game.GetWidth(), game.GetHeight(), game.GetMines()}
		_ = flow.GetRouterInstance().RouteTo(Routing.RouteGame, params)
	} else {
		_ = flow.GetRouterInstance().RouteTo(Routing.RouteCustomGame, nil)
	}
}

func (c *Controller) String() string {
	return "[go_mines/SelectGame.Controller]()"
}

func (c *Controller) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] %v Does not understand: %v",
		time.Now().Format(time.UnixDate),
		c,
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" construction function
// ---------------------------------------------------------------------------------------------------------------------

func NewController() *Controller {
	return &Controller{}
}
