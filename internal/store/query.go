package store

import (
	"context"
	"database/sql"
	"strings"
)

func (s *Store) SearchSongIDs(ctx context.Context, q string) ([]int64, error) {
	rows, e := s.db.QueryContext(ctx, `SELECT id FROM songs WHERE lower(title) LIKE ? OR lower(composer) LIKE ?`, "%"+strings.ToLower(q)+"%", "%"+strings.ToLower(q)+"%")
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []int64{}
	for rows.Next() {
		var id int64
		if e = rows.Scan(&id); e != nil {
			return nil, e
		}
		out = append(out, id)
	}
	return out, rows.Err()
}
func (s *Store) WithRows(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return s.db.QueryContext(ctx, query, args...)
}
func (s *Store) UpdateSongLanguage(ctx context.Context, id int64, language string) error {
	_, e := s.db.ExecContext(ctx, `UPDATE songs SET language=? WHERE id=?`, language, id)
	return e
}
func (s *Store) UpdateScoreFeatured(ctx context.Context, id int64, featured bool) error {
	_, e := s.db.ExecContext(ctx, `UPDATE scores SET featured=? WHERE id=?`, featured, id)
	return e
}
func (s *Store) DeleteScores(ctx context.Context, songID int64) error {
	_, e := s.db.ExecContext(ctx, `DELETE FROM scores WHERE song_id=?`, songID)
	return e
}
