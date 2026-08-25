package store

import (
	"context"
	"fmt"
	"time"
)

func (s *Store) TableRows(ctx context.Context, table string) (int, error) {
	var n int
	if e := s.db.QueryRowContext(ctx, "SELECT count(*) FROM "+table).Scan(&n); e != nil {
		return 0, e
	}
	return n, nil
}
func (s *Store) DatabaseInfo(ctx context.Context) (string, error) {
	var v string
	e := s.db.QueryRowContext(ctx, "SELECT sqlite_version()").Scan(&v)
	return v, e
}
func (s *Store) Integrity(ctx context.Context) (bool, error) {
	var v string
	e := s.db.QueryRowContext(ctx, "PRAGMA integrity_check").Scan(&v)
	return v == "ok", e
}
func (s *Store) CreatedAt(ctx context.Context, id int64) (time.Time, error) {
	x, e := s.GetSong(ctx, id)
	return x.CreatedAt, e
}
func (s *Store) Describe(ctx context.Context) string {
	st, e := s.Stats(ctx)
	if e != nil {
		return "error"
	}
	return fmt.Sprintf("songs=%d scores=%d members=%d rehearsals=%d", st.Songs, st.Scores, st.Members, st.Rehearsals)
}
