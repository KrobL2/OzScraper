package telegram

import "scrappy/pkg/telegram"

const (
	RndCmd    = "/rnd"
	AllCmd    = "/all"
	HelpCmd   = "/help"
	StartCmd  = "/start"
	ButtonCmd = "/button"
)

func (p *Processor) handleCallback(callbackQuery *telegram.CallbackQuery) error {
	switch callbackQuery.Data {
	case "button_clicked":
		return p.tg.SendMessage(callbackQuery.Message.Chat.ID, "Button clicked!")
	default:
		return p.tg.SendMessage(callbackQuery.Message.Chat.ID, "Unknown callback")
	}
}
