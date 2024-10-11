package files

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
)

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save page", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}
