package files

import "read-adviser-bot/storage"

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
