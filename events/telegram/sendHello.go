package telegram

func (p *Processor) sendHello(chatID int) error {
	err := p.sendButtonMessage(chatID)
	if err != nil {
		return err
	}

	return p.tg.SendMessage(chatID, msgHello)
}

/* func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}
*/
