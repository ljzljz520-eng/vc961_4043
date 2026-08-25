package store

import (
	"choirsearch/internal/model"
	"context"
)

func (s *Store) SaveScore(ctx context.Context, x model.Score) error {
	_, e := s.db.ExecContext(ctx, `INSERT OR REPLACE INTO scores(id,song_id,voice,format,uri,pages,featured) VALUES(?,?,?,?,?,?,?)`, x.ID, x.SongID, x.Voice, x.Format, x.URI, x.Pages, x.Featured)
	return e
}
func (s *Store) ScoresForSong(ctx context.Context, id int64) ([]model.Score, error) {
	rows, e := s.db.QueryContext(ctx, `SELECT id,song_id,voice,format,uri,pages,featured FROM scores WHERE song_id=?`, id)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Score{}
	for rows.Next() {
		var x model.Score
		if e = rows.Scan(&x.ID, &x.SongID, &x.Voice, &x.Format, &x.URI, &x.Pages, &x.Featured); e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func (s *Store) ScoresByVoice(ctx context.Context, v string) ([]model.Score, error) {
	rows, e := s.db.QueryContext(ctx, `SELECT id,song_id,voice,format,uri,pages,featured FROM scores WHERE voice=?`, v)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []model.Score{}
	for rows.Next() {
		var x model.Score
		if e = rows.Scan(&x.ID, &x.SongID, &x.Voice, &x.Format, &x.URI, &x.Pages, &x.Featured); e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func (s *Store) ScoreCount(ctx context.Context, id int64) (int, error) {
	var n int
	e := s.db.QueryRowContext(ctx, "SELECT count(*) FROM scores WHERE song_id=?", id).Scan(&n)
	return n, e
}
