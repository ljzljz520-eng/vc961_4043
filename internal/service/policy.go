package service

import (
	"choirsearch/internal/model"
	"context"
	"fmt"
	"strings"
)

type Policy struct {
	MaxResults     int
	AllowedFormats map[string]bool
	RequiredVoices []string
}

func DefaultPolicy() Policy {
	return Policy{MaxResults: 100, AllowedFormats: map[string]bool{"pdf": true, "mxl": true, "musicxml": true}, RequiredVoices: []string{"S", "A", "T", "B"}}
}
func (p Policy) ValidateScore(x model.Score) error {
	if x.SongID <= 0 {
		return fmt.Errorf("song required")
	}
	if x.Pages < 1 {
		return fmt.Errorf("pages required")
	}
	if len(p.AllowedFormats) > 0 && !p.AllowedFormats[strings.ToLower(x.Format)] {
		return fmt.Errorf("format not allowed")
	}
	return nil
}
func (p Policy) ValidateVoice(v string) bool {
	for _, x := range p.RequiredVoices {
		if x == v {
			return true
		}
	}
	return false
}
func (p Policy) Limit(n int) int {
	if n <= 0 {
		return p.MaxResults
	}
	if n > p.MaxResults {
		return p.MaxResults
	}
	return n
}
func (s *Service) RegisterValidated(ctx context.Context, x model.Song, ss []model.Score, p Policy) error {
	if !x.IsValid() {
		return fmt.Errorf("invalid song")
	}
	for _, z := range ss {
		if e := p.ValidateScore(z); e != nil {
			return e
		}
	}
	return s.RegisterSong(ctx, x, ss)
}
func (s *Service) RemoveSong(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid id")
	}
	return s.catalog.Remove(ctx, id)
}
func (s *Service) RenameLanguage(ctx context.Context, id int64, language string) error {
	if strings.TrimSpace(language) == "" {
		return fmt.Errorf("language required")
	}
	return s.store.UpdateSongLanguage(ctx, id, language)
}
func (s *Service) FeatureScore(ctx context.Context, id int64, featured bool) error {
	return s.store.UpdateScoreFeatured(ctx, id, featured)
}
func (s *Service) PurgeSong(ctx context.Context, id int64) error {
	if e := s.store.DeleteScores(ctx, id); e != nil {
		return e
	}
	return s.RemoveSong(ctx, id)
}
