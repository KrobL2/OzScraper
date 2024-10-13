package events

type Event struct {
	Type          EventType
	Text          string
	Meta          interface{}
	CallbackQuery interface{}
}

type EventType int

const (
	Message EventType = iota
	CallbackQuery
)
