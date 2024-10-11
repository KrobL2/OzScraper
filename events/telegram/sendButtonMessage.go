package telegram

import "read-adviser-bot/clients/telegram"

func (p *Processor) sendButtonMessage(chatID int) error {
	keyboard := &telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			{
				{
					Text:         "Click me!",
					CallbackData: "button_clicked",
				},
			},
		},
	}

	return p.tg.SendMessageWithKeyboard(chatID, "Here's a button:", keyboard)
}
