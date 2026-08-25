package store

import (
	"choirsearch/internal/model"
	"context"
	"time"
)

func (s *Store) SaveMember(ctx context.Context, x model.Member) error {
	_, e := s.db.ExecContext(ctx, `INSERT OR REPLACE INTO members(id,name,email,voice,joined_at,active) VALUES(?,?,?,?,?,?)`, x.ID, x.Name, x.Email, x.Voice, x.JoinedAt.Format(time.RFC3339), x.Active)
	return e
}
func (s *Store) ActiveMembers(ctx context.Context) ([]model.Member, error) {
	rows, e := s.db.QueryContext(ctx, `SELECT id,name,email,voice,joined_at,active FROM members WHERE active=1`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Member{}
	for rows.Next() {
		var x model.Member
		var st string
		if e = rows.Scan(&x.ID, &x.Name, &x.Email, &x.Voice, &st, &x.Active); e != nil {
			return nil, e
		}
		x.JoinedAt, _ = time.Parse(time.RFC3339, st)
		out = append(out, x)
	}
	return out, rows.Err()
}
