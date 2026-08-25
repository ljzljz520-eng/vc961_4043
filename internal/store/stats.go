package store

import (
	"context"
	"database/sql"
)

type Stats struct{ Songs, Scores, Members, Rehearsals int }

func (s *Store) Stats(ctx context.Context) (Stats, error) {
	var r Stats
	queries := []*int{&r.Songs, &r.Scores, &r.Members, &r.Rehearsals}
	names := []string{"songs", "scores", "members", "rehearsals"}
	for i, n := range names {
		if e := s.db.QueryRowContext(ctx, "SELECT count(*) FROM "+n).Scan(queries[i]); e != nil {
			return r, e
		}
	}
	return r, nil
}
func (s *Store) Vacuum(ctx context.Context) error { _, e := s.db.ExecContext(ctx, "VACUUM"); return e }
func (s *Store) Tx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, e := s.db.BeginTx(ctx, nil)
	if e != nil {
		return e
	}
	if e = fn(tx); e != nil {
		tx.Rollback()
		return e
	}
	return tx.Commit()
}
