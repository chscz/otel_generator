package spanaction

type EventAttribute interface{}

type Event struct{}

func NewEvent(attrGenerator EventAttribute) *Event {
	return &Event{}
}
