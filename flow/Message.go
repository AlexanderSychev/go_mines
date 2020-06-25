package flow

import (
	"fmt"
	"reflect"
	"strings"
)

// ---------------------------------------------------------------------------------------------------------------------
// "NilHandleError" error definition
// ---------------------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------------------
// "Message" type definition
// ---------------------------------------------------------------------------------------------------------------------

// "Message" - object which provides indirection calls of method of the other entity
type Message struct {
	// Name of target object method
	method string
	// Arguments for target object method
	arguments []reflect.Value
}

// ---------------------------------------------------------------------------------------------------------------------
// "Message" type public methods
// ---------------------------------------------------------------------------------------------------------------------

// Stringify message (for "fmt")
func (m Message) String() string {
	strArguments := make([]string, len(m.arguments))
	for index, value := range m.arguments {
		strArguments[index] = fmt.Sprintf("%v", value)
	}
	return fmt.Sprintf("MESSAGE::%s(%s)", m.method, strings.Join(strArguments, ","))
}

// Add arbitrary type argument to the end
func (m *Message) PushArg(arg interface{}) {
	m.arguments = append(m.arguments, reflect.ValueOf(arg))
}

// Add arbitrary type argument to the start
func (m *Message) UnshiftArg(arg interface{}) {
	toAdd := make([]reflect.Value, 1)
	toAdd[0] = reflect.ValueOf(arg)
	m.arguments = append(toAdd, m.arguments...)
}

// ---------------------------------------------------------------------------------------------------------------------
// "Message" type construction functions
// ---------------------------------------------------------------------------------------------------------------------

// Creates new "Message" instance with empty arguments list
func NewMessage(method string) Message {
	return Message{
		method:    method,
		arguments: make([]reflect.Value, 0),
	}
}

// Creates new "Message" instance with arguments
func NewMessageWithArgs(method string, args ...interface{}) Message {
	arguments := make([]reflect.Value, len(args))
	for index, arg := range args {
		arguments[index] = reflect.ValueOf(arg)
	}

	return Message{
		method:    method,
		arguments: arguments,
	}
}
