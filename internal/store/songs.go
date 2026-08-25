package store

import (
	"choirsearch/internal/model"
	"context"
	"strings"
	"time"
)

func (s *Store) SaveSong(ctx context.Context, x model.Song) error {
	_, e := s.db.ExecContext(ctx, `INSERT OR REPLACE INTO songs(id,title,composer,language,difficulty,voices,tags,created_at) VALUES(?,?,?,?,?,?,?,?)`, x.ID, x.Title, x.Composer, x.Language, x.Difficulty, strings.Join(x.Voices, ","), strings.Join(x.Tags, ","), x.CreatedAt.Format(time.RFC3339))
	return e
}
func (s *Store) GetSong(ctx context.Context, id int64) (model.Song, error) {
	var x model.Song
	var v, t, created string
	e := s.db.QueryRowContext(ctx, `SELECT id,title,composer,language,difficulty,voices,tags,created_at FROM songs WHERE id=?`, id).Scan(&x.ID, &x.Title, &x.Composer, &x.Language, &x.Difficulty, &v, &t, &created)
	if e != nil {
		return x, e
	}
	x.Voices = strings.Split(v, ",")
	x.Tags = strings.Split(t, ",")
	x.CreatedAt, _ = time.Parse(time.RFC3339, created)
	return x, nil
}
func (s *Store) ListSongs(ctx context.Context) ([]model.Song, error) {
	rows, e := s.db.QueryContext(ctx, `SELECT id FROM songs ORDER BY title`)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Song{}
	for rows.Next() {
		var id int64
		if e = rows.Scan(&id); e != nil {
			return nil, e
		}
		x, e := s.GetSong(ctx, id)
		if e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func (s *Store) DeleteSong(ctx context.Context, id int64) error {
	_, e := s.db.ExecContext(ctx, `DELETE FROM songs WHERE id=?`, id)
	return e
}
func (s *Store) CountSongs(ctx context.Context) (int, error) {
	var n int
	e := s.db.QueryRowContext(ctx, `SELECT count(*) FROM songs`).Scan(&n)
	return n, e
}
