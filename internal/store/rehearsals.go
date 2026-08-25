package store

import (
	"choirsearch/internal/model"
	"context"
	"time"
)

func (s *Store) SaveRehearsal(ctx context.Context, x model.Rehearsal) error {
	_, e := s.db.ExecContext(ctx, `INSERT OR REPLACE INTO rehearsals(id,name,venue,starts_at,duration,conductor) VALUES(?,?,?,?,?,?)`, x.ID, x.Name, x.Venue, x.StartsAt.Format(time.RFC3339), x.Duration, x.Conductor)
	return e
}
func (s *Store) ListRehearsals(ctx context.Context) ([]model.Rehearsal, error) {
	rows, e := s.db.QueryContext(ctx, `SELECT id,name,venue,starts_at,duration,conductor FROM rehearsals ORDER BY starts_at`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Rehearsal{}
	for rows.Next() {
		var x model.Rehearsal
		var st string
		if e = rows.Scan(&x.ID, &x.Name, &x.Venue, &st, &x.Duration, &x.Conductor); e != nil {
			return nil, e
		}
		x.StartsAt, _ = time.Parse(time.RFC3339, st)
		out = append(out, x)
	}
	return out, rows.Err()
}
