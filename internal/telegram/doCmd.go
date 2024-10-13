package telegram

import (
	"log"
	"strings"
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
	case AllCmd:
		return p.sendAll(chatID, username)
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
