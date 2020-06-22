package pubsub

// Subscriber can handle events from Publisher - it's just must implement
// "HandleEvent" method
type Subscriber interface {
	HandleEvent(event Event)
}
