package telegram

import (
	"log"
	"strings"

	"read-adviser-bot/clients/telegram"
)

const (
	RndCmd    = "/rnd"
	HelpCmd   = "/help"
	StartCmd  = "/start"
	ButtonCmd = "/button"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	if IsAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	case ButtonCmd:
		return p.sendButtonMessage(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) handleCallback(callbackQuery *telegram.CallbackQuery) error {
	switch callbackQuery.Data {
	case "button_clicked":
		return p.tg.SendMessage(callbackQuery.Message.Chat.ID, "Button clicked!")
	default:
		return p.tg.SendMessage(callbackQuery.Message.Chat.ID, "Unknown callback")
	}
}
