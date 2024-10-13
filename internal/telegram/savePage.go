package telegram

import (
	"context"
	"scrappy/internal/storage"
	e "scrappy/pkg/errors"
)

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() {
		err = e.WrapIfErr("can't do command: save page", err)
	}()

	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(context.Background(), page)
	if err != nil {
		return err
	}

	if isExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(context.Background(), page); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}
