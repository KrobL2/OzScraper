package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"read-adviser-bot/storage"
)

// GetAll pages from storage.
func (s *Storage) GetAll(ctx context.Context, userName string) (*storage.Page, error) {
	q := `SELECT * FROM pages`

	var url string

	err := s.db.QueryRowContext(ctx, q, userName).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, storage.ErrNoSavedPages
	}

	if err != nil {
		return nil, fmt.Errorf("can't get pages: %w", err)
	}

	return &storage.Page{
		URL:      url,
		UserName: userName,
	}, nil
}
