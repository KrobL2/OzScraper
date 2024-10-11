package sqlite

import (
	"context"
	"fmt"
	"read-adviser-bot/storage"
)

// IsExists checks if page exists in storage.
func (s *Storage) IsExists(ctx context.Context, page *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`

	var count int

	if err := s.db.QueryRowContext(ctx, q, page.URL, page.UserName).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check if page exists: %w", err)
	}

	return count > 0, nil
}