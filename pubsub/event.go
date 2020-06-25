package pubsub

// ---------------------------------------------------------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------------------------------------------------------

const (
	CellOpenedEvent           = 0
	CellToggleMarkEvent       = 1
	CellButtonLeftClickEvent  = 2
	CellButtonRightClickEvent = 3
	CreateNewFieldEvent       = 4
	CustomGameEvent           = 5
)

// ---------------------------------------------------------------------------------------------------------------------
// "Event" type definition
// ---------------------------------------------------------------------------------------------------------------------

type Event struct {
	Type    uint16
	Payload interface{}
}

func NewEvent(eventType uint16, payload interface{}) Event {
	return Event{eventType, payload}
}

// Concrete Events

type CellOpenedEventPayload struct {
	X, Y        int
	HasMine     bool
	IsMarked    bool
	MinesAround int
}

func NewCellOpenedEvent(payload CellOpenedEventPayload) Event {
	return NewEvent(CellOpenedEvent, payload)
}

type CellToggleMarkEventPayload struct {
	X, Y     int
	IsMarked bool
}

func NewCellToggleMarkEvent(payload CellToggleMarkEventPayload) Event {
	return NewEvent(CellToggleMarkEvent, payload)
}

type CellButtonClickEventPayload struct {
	X, Y int
}

func NewCellButtonClickEvent(right bool, payload CellButtonClickEventPayload) Event {
	var eventType uint16 = CellButtonLeftClickEvent
	if right {
		eventType = CellButtonRightClickEvent
	}
	return NewEvent(eventType, payload)
}

type CreateNewFieldEventPayload struct {
	width, height, mines int
}

func NewCreateNewFieldEvent(width, height, mines int) Event {
	return NewEvent(CreateNewFieldEvent, CreateNewFieldEventPayload{width, height, mines})
}

func NewCustomGameEvent() Event {
	return NewEvent(CustomGameEvent, nil)
}
