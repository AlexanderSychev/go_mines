package flow

import (
	"fmt"
	"reflect"
)

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	// Reserved key for application actor
	ApplicationActor = "application"
	// Reserved key for application router actor
	RouterActor = "router"
	// Default key for controller actor
	ControllerActor = "controller"
	// Default key for view actor
	ViewActor = "view"
	// Default key for model actor
	ModelActor = "model"
)

var restrictedMethods = [...]string{
	"HandleUnknownMessage",
	"GetWidgets",
	"Activate",
	"Deactivate",
}

// ---------------------------------------------------------------------------------------------------------------------
// "Broker" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Broker struct {
	actors map[string]IActor
}

// ---------------------------------------------------------------------------------------------------------------------
// "Broker" type methods
// ---------------------------------------------------------------------------------------------------------------------

func (broker *Broker) isRestrictedMethod(message Message) bool {
	for _, method := range restrictedMethods {
		if message.method == method {
			return true
		}
	}
	return false
}

func (broker *Broker) RouteTo(route string, params interface{}) ([]reflect.Value, error) {
	return broker.SendToActor(RouterActor, NewMessageWithArgs("RouteTo", route, params))
}

func (broker *Broker) SendToActor(actorName string, message Message) ([]reflect.Value, error) {
	if broker.isRestrictedMethod(message) {
		return nil, NewError(
			"Restricted Method",
			fmt.Sprintf("Method \"%s\" is restricted to use", message.method),
		)
	}

	actor, actorExists := broker.actors[actorName]
	if !actorExists {
		return nil, NewError("Unknown Actor", fmt.Sprintf("Actor \"%s\" not found", actorName))
	}

	vActor := reflect.ValueOf(actor)
	if vActor.IsValid() {
		vMethod := vActor.MethodByName(message.method)
		if vMethod.IsValid() {
			return vMethod.Call(message.arguments), nil
		} else {
			return actor.HandleUnknownMessage(message)
		}
	} else {
		return nil, NewError(
			"Invalid Actor",
			fmt.Sprintf("Can't handle message %v by actor (%v, %T)", message, actor, actor),
		)
	}
}

func (broker *Broker) AddActor(actorName string, actor IActor) {
	broker.actors[actorName] = actor
}

func (broker *Broker) RemoveActor(actorName string) {
	delete(broker.actors, actorName)
}

// ---------------------------------------------------------------------------------------------------------------------
// "Broker" type lazy singleton implementation
// ---------------------------------------------------------------------------------------------------------------------

// Global "Broker" instance pointer. DO NOT USE IT!
var broker *Broker = nil

// Returns global "Broker" instance pointer. If global "Broker" not created yet, this function creates it.
func GetBrokerInstance() *Broker {
	if broker == nil {
		broker = &Broker{
			actors: make(map[string]IActor),
		}
	}
	return broker
}
