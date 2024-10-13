package telegram

import "context"

func (p *Processor) sendAll(chatID int, username string) error {
	pages, err := p.storage.GetAll(context.Background(), username)
	if err != nil {
		return err
	}

	if len(pages) == 0 {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	for _, page := range pages {
		err := p.tg.SendMessage(chatID, page.URL)
		if err != nil {
			return err
		}
	}

	return nil
}
