package flow

import "reflect"

// ---------------------------------------------------------------------------------------------------------------------
// "IActor" interface definition
// ---------------------------------------------------------------------------------------------------------------------

// Common interface for any entities what can receive "Message" from "Broker"
type IActor interface {
	// Special method to handle unknown message. This can be useful for logging or something third party actions.
	// Don't use it as "universal" message handler.
	// Only FATAL error should be returned (application will fail with panic), otherwise should return "nil"
	HandleUnknownMessage(message Message) ([]reflect.Value, error)
}
