package pubsub

// -----------------------------------------------------------------------------
// Constants
// -----------------------------------------------------------------------------

const (
	CELL_OPENED_EVENT             = 0
	CELL_TOGGLE_MARK_EVENT        = 1
	CELL_BUTTON_LEFT_CLICK_EVENT  = 2
	CELL_BUTTON_RIGHT_CLICK_EVENT = 3
)

// -----------------------------------------------------------------------------
// "Event" type definition
// -----------------------------------------------------------------------------

type Event struct {
	Type    byte
	Payload interface{}
}

func NewEvent(eventType byte, payload interface{}) Event {
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
	return NewEvent(CELL_OPENED_EVENT, payload)
}

type CellToggleMarkEventPayload struct {
	X, Y     int
	IsMarked bool
}

func NewCellToggleMarkEvent(payload CellToggleMarkEventPayload) Event {
	return NewEvent(CELL_TOGGLE_MARK_EVENT, payload)
}

type CellButtonClickEventPayload struct {
	X, Y int
}

func NewCellButtonClickEvent(
	right bool,
	payload CellButtonClickEventPayload,
) Event {
	var eventType byte = CELL_BUTTON_LEFT_CLICK_EVENT
	if right {
		eventType = CELL_BUTTON_RIGHT_CLICK_EVENT
	}
	return NewEvent(eventType, payload)
}
