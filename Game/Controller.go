package Game

import (
	"github.com/AlexanderSychev/go_mines/flow"
	"log"
	"math"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Controller struct {}

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (c *Controller) openEmptyNotMarkedNeighbors(x, y int) {
	broker := flow.GetBrokerInstance()
	res, err := broker.SendToActor(flow.ModelActor, flow.NewMessageWithArgs("GetEmptyNotMarkedNeighbors", x, y))
	if err != nil {
		log.Println(err)
		return
	}

	points := res[0].Interface().([]Point)
	for _, point := range points {
		res, err = broker.SendToActor(flow.ModelActor, flow.NewMessageWithArgs("Open", point.X(), point.Y()))
		if err != nil {
			log.Println(err)
			return
		}
		minesAround := int(res[1].Int())
		_, err = broker.SendToActor(flow.ViewActor, flow.NewMessageWithArgs(
			"SetOpened",
			point.X(),
			point.Y(),
			minesAround,
		))
	}
}

func (c *Controller) openAllWithMines() {
	broker := flow.GetBrokerInstance()
	res, err := broker.SendToActor(flow.ModelActor, flow.NewMessage("GetAllWithMines"))
	if err != nil {
		log.Println(err)
		return
	}

	points := res[0].Interface().([]Point)
	for _, point := range points {
		_, err = broker.SendToActor(flow.ModelActor, flow.NewMessageWithArgs("Open", point.X(), point.Y()))
		if err != nil {
			log.Println(err)
			return
		}
		_, err = broker.SendToActor(flow.ViewActor, flow.NewMessageWithArgs(
			"SetWithMine",
			point.X(),
			point.Y(),
		))
	}
}

func (c *Controller) UpdateTime(d time.Duration) {
	minutes := int(math.Floor(d.Minutes()))
	seconds := 60 * minutes - int(math.Floor(d.Seconds()))
	message := flow.NewMessageWithArgs("SetClock", minutes, seconds)
	_, err := flow.GetBrokerInstance().SendToActor(flow.ViewActor, message)
	if err != nil {
		log.Println(err)
	}
}

func (c *Controller) Timeout() {
	_, err := flow.GetBrokerInstance().SendToActor(flow.ViewActor, flow.NewMessage("GameOver"))
	if err != nil {
		log.Println(err)
	}
}

func (c *Controller) OnMark(x, y int) {
	modelMessage := flow.NewMessageWithArgs("ToggleMark", x, y)

	res, err := flow.GetBrokerInstance().SendToActor(flow.ModelActor, modelMessage)
	if err != nil {
		log.Println(err)
		return
	}

	isMarked := res[0].Bool()
	viewMethod := "SetClosed"
	if isMarked {
		viewMethod = "SetMarked"
	}
	viewMessage := flow.NewMessageWithArgs(viewMethod, x, y)
	_, err = flow.GetBrokerInstance().SendToActor(flow.ViewActor, viewMessage)
}

func (c *Controller) OnOpen(x, y int) {
	broker := flow.GetBrokerInstance()
	modelMessage := flow.NewMessageWithArgs("Open", x, y)

	res, err := broker.SendToActor(flow.ModelActor, modelMessage)
	if err != nil {
		log.Println(err)
		return
	}

	hasMine := res[0].Bool()
	minesAround := int(res[1].Int())

	if hasMine {
		c.openAllWithMines()
		_, err = broker.SendToActor(flow.ViewActor, flow.NewMessage("GameOver"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	_, err = broker.SendToActor(flow.ViewActor, flow.NewMessageWithArgs(
		"SetOpened",
		x,
		y,
		minesAround,
	))

	if minesAround == 0 {
		c.openEmptyNotMarkedNeighbors(x, y)
	}
}

func (c *Controller) HandleUnknownMessage(message flow.Message) ([]reflect.Value, error) {
	log.Printf(
		"[%s] [go_mines/Game.Controller]() Does not understand: %v",
		time.Now().Format(time.UnixDate),
		message,
	)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Controller" type constructor
// ---------------------------------------------------------------------------------------------------------------------

func NewController() *Controller {
	return &Controller{}
}
