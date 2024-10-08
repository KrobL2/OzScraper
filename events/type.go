package events

import "read-adviser-bot/clients/telegram"

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
	CallbackQuery = 2
)

type Event struct {
	Type          Type
	Text          string
	Meta          interface{}
	CallbackQuery *telegram.CallbackQuery
}
