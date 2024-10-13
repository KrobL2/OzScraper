package sqlite

import (
	"context"
	"fmt"
	"scrappy/storage"
)

// GetAll pages from storage.
func (s *Storage) GetAll(ctx context.Context, userName string) ([]storage.Page, error) {
	q := `SELECT url FROM pages WHERE user_name = ?`

	rows, err := s.db.QueryContext(ctx, q, userName)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить страницы: %w", err)
	}
	defer rows.Close()

	var pages []storage.Page
	for rows.Next() {
		var p storage.Page
		if err := rows.Scan(&p.URL); err != nil {
			return nil, fmt.Errorf("не удалось сканировать страницу: %w", err)
		}
		p.UserName = userName
		pages = append(pages, p)
	}

	return pages, nil
}
