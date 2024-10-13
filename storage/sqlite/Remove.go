package sqlite

import (
	"context"
	"fmt"
	"scrappy/storage"
)

// Remove removes page from storage.
func (s *Storage) Remove(ctx context.Context, page *storage.Page) error {
	q := `DELETE FROM pages WHERE url = ? AND user_name = ?`

	if _, err := s.db.ExecContext(ctx, q, page.URL, page.UserName); err != nil {
		return fmt.Errorf("can't remove page: %w", err)
	}

	return nil
}
