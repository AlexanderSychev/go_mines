package flow

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

// ---------------------------------------------------------------------------------------------------------------------
// "Router" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Router struct {
	routes map[string]PageCreator
}

// ---------------------------------------------------------------------------------------------------------------------
// "Router" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (router *Router) RouteTo(route string, params interface{}) error {
	creator, ok := router.routes[route]
	if ok {
		page, err := creator(params)
		if err != nil {
			return err
		}
		_, err = GetBrokerInstance().SendToActor(ApplicationActor, NewMessageWithArgs("OnRoute", page))
		if err != nil {
			return err
		}
		return nil
	} else {
		return NewError("Unknown Route", fmt.Sprintf("Route %s is undefined", route))
	}
}

func (router *Router) AddRoute(route string, creator PageCreator) {
	router.routes[route] = creator
}

func (router *Router) RemoveRoute(route string) {
	delete(router.routes, route)
}

func (router *Router) HandleUnknownMessage(message Message) ([]reflect.Value, error) {
	log.Printf("[go_mines/flow.Router] [%s] Does not understand: %v", time.Now().Format(time.UnixDate), message)
	return nil, nil
}

// ---------------------------------------------------------------------------------------------------------------------
// "Router" type lazy singleton implementation
// ---------------------------------------------------------------------------------------------------------------------

var router *Router = nil

func GetRouterInstance() *Router {
	if router == nil {
		router = &Router{
			routes: make(map[string]PageCreator),
		}
		GetBrokerInstance().AddActor(RouterActor, router)
	}
	return router
}
