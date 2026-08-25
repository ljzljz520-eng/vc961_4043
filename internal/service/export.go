package service

import (
	"choirsearch/internal/model"
	"context"
	"fmt"
	"strings"
)

func (s *Service) ExportAll(ctx context.Context) (string, error) {
	xs, e := s.store.ListSongs(ctx)
	if e != nil {
		return "", e
	}
	lines := []string{}
	for _, x := range xs {
		lines = append(lines, model.FormatSongLine(x))
	}
	return strings.Join(lines, "\n"), nil
}
func (s *Service) ExportScores(ctx context.Context, id int64) (string, error) {
	xs, e := s.store.ScoresForSong(ctx, id)
	if e != nil {
		return "", e
	}
	lines := []string{}
	for _, x := range xs {
		lines = append(lines, fmt.Sprintf("%d|%s|%s|%s|%d", x.ID, x.Voice, x.Format, x.URI, x.Pages))
	}
	return strings.Join(lines, "\n"), nil
}
func (s *Service) ImportScores(ctx context.Context, id int64, lines []string) error {
	for i, line := range lines {
		p := strings.Split(line, "|")
		if len(p) < 3 {
			return fmt.Errorf("score line %d", i)
		}
		x := model.Score{ID: int64(i + 1), SongID: id, Voice: p[0], Format: p[1], URI: p[2]}
		if e := s.store.SaveScore(ctx, x); e != nil {
			return e
		}
	}
	return nil
}
func (s *Service) SongSummary(ctx context.Context, id int64) (string, error) {
	x, e := s.store.GetSong(ctx, id)
	if e != nil {
		return "", e
	}
	n, e := s.store.ScoreCount(ctx, id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%s by %s (%d scores)", x.Title, x.Composer, n), nil
}
