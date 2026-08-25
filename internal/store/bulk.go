package store

import (
	"choirsearch/internal/model"
	"context"
	"strings"
	"time"
)

func (s *Store) SaveSongs(ctx context.Context, xs []model.Song) error {
	for _, x := range xs {
		if e := s.SaveSong(ctx, x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) ReplaceSongs(ctx context.Context, xs []model.Song) error {
	for _, x := range xs {
		if e := s.SaveSong(ctx, x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SaveScores(ctx context.Context, xs []model.Score) error {
	for _, x := range xs {
		if e := s.SaveScore(ctx, x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SaveRehearsals(ctx context.Context, xs []model.Rehearsal) error {
	for _, x := range xs {
		if e := s.SaveRehearsal(ctx, x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SaveMembers(ctx context.Context, xs []model.Member) error {
	for _, x := range xs {
		if e := s.SaveMember(ctx, x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) SongExists(ctx context.Context, id int64) bool {
	_, e := s.GetSong(ctx, id)
	return e == nil
}
func (s *Store) ScoreExists(ctx context.Context, id int64) bool {
	var n int
	return s.db.QueryRowContext(ctx, "SELECT count(*) FROM scores WHERE id=?", id).Scan(&n) == nil && n > 0
}
func (s *Store) MemberCount(ctx context.Context) int {
	var n int
	s.db.QueryRowContext(ctx, "SELECT count(*) FROM members").Scan(&n)
	return n
}
func (s *Store) RehearsalCount(ctx context.Context) int {
	var n int
	s.db.QueryRowContext(ctx, "SELECT count(*) FROM rehearsals").Scan(&n)
	return n
}
func (s *Store) TouchSong(ctx context.Context, id int64) error {
	return s.UpdateSongLanguage(ctx, id, strings.TrimSpace(time.Now().Format("2006")))
}
