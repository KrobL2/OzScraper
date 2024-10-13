package telegram

import (
	"scrappy/clients/telegram"
)

func (p *Processor) sendButtonMessage(chatID int) error {
	commands := []telegram.BotCommand{
		{Command: "/start", Description: "Начать работу с ботом"},
		{Command: "/help", Description: "Получить помощь"},
		{Command: "/rnd", Description: "Получить случайную ссылку"},
		{Command: "/all", Description: "Получить все сохраненные ссылки"},
	}

	err := p.tg.SetBotCommands(commands)
	if err != nil {
		return err
	}

	return p.tg.SendMessage(chatID, "Команды бота обновлены. Проверьте меню команд.")
}

/* func (p *Processor) sendButtonMessage(chatID int) error {
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
*/
