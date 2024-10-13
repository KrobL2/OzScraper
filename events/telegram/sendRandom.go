package telegram

import (
	"context"
	"errors"
	"scrappy/lib/e"
	"scrappy/storage"
)

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() {
		err = e.WrapIfErr("can't do command: can't send random", err)
	}()

	page, err := p.storage.PickRandom(context.Background(), username)
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
