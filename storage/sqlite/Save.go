package sqlite

import (
	"context"
	"fmt"
	"read-adviser-bot/storage"
)

// Save saves page to storage.
func (s *Storage) Save(ctx context.Context, p *storage.Page) error {
	q := `INSERT INTO pages (url, user_name) VALUES (?, ?)`

	if _, err := s.db.ExecContext(ctx, q, p.URL, p.UserName); err != nil {
		return fmt.Errorf("can't save page: %w", err)
	}

	return nil
}
