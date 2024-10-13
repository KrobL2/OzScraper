package telegram

import (
	"context"
	"errors"
	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
)

func (p *Processor) sendAll(chatID int, username string) (err error) {
	defer func() {
		err = e.WrapIfErr("can't do command: can't send all", err)
	}()

	page, err := p.storage.GetAll(context.Background(), username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}

	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return err
	}

	return p.storage.Remove(context.Background(), page)
}
