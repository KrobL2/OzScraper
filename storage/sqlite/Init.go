package sqlite

import (
	"context"
	"fmt"
)

func (s *Storage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
