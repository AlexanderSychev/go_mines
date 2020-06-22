package pubsub

// -----------------------------------------------------------------------------
// "Publisher" type definition
// -----------------------------------------------------------------------------

type Publisher struct {
	channels map[byte][]*Subscriber
}

// Methods

func (pub *Publisher) Subscribe(eventType byte, sub *Subscriber) {
	if pub.channels[eventType] == nil {
		pub.channels[eventType] = make([]*Subscriber, 0)
	}
	pub.channels[eventType] = append(pub.channels[eventType], sub)
}

func (pub *Publisher) Unsubscribe(eventType byte, sub *Subscriber) {
	if pub.channels[eventType] != nil {
		newList := make([]*Subscriber, 0)
		for _, s := range pub.channels[eventType] {
			if s != sub {
				newList = append(newList, s)
			}
		}
		pub.channels[eventType] = newList
	}
}

func (pub *Publisher) Publish(event Event) {
	if pub.channels[event.Type] != nil {
		for _, sub := range pub.channels[event.Type] {
			if sub != nil {
				(*sub).HandleEvent(event)
			}
		}
	}
}

func (pub *Publisher) BulkPublish(events []Event) {
	if events != nil && len(events) > 0 {
		for _, event := range events {
			pub.Publish(event)
		}
	}
}

// Constructor

func NewPublisher() *Publisher {
	return &Publisher{make(map[byte][]*Subscriber)}
}
