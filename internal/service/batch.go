package service

import (
	"choirsearch/internal/model"
	"context"
	"fmt"
)

func (s *Service) BulkRegister(ctx context.Context, xs []model.Song) (int, error) {
	n := 0
	for _, x := range xs {
		if e := s.RegisterSong(ctx, x, nil); e != nil {
			return n, e
		}
		n++
	}
	return n, nil
}
func (s *Service) ValidateSongs(xs []model.Song) []error {
	out := []error{}
	for _, x := range xs {
		if !x.IsValid() {
			out = append(out, fmt.Errorf("song %d invalid", x.ID))
		}
	}
	return out
}
func (s *Service) SearchAll(ctx context.Context, voices []string) (map[string][]model.SearchResult, error) {
	out := map[string][]model.SearchResult{}
	for _, v := range voices {
		r, e := s.FindScores(ctx, model.SearchQuery{Voice: v})
		if e != nil {
			return out, e
		}
		out[v] = r
	}
	return out, nil
}
func (s *Service) SearchFirst(ctx context.Context, voices []string) ([]model.SearchResult, error) {
	for _, v := range voices {
		r, e := s.FindScores(ctx, model.SearchQuery{Voice: v})
		if e != nil {
			return nil, e
		}
		if len(r) > 0 {
			return r, nil
		}
	}
	return []model.SearchResult{}, nil
}
func (s *Service) DeleteMany(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		if e := s.RemoveSong(ctx, id); e != nil {
			return e
		}
	}
	return nil
}
func (s *Service) SetFeatured(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		if e := s.FeatureScore(ctx, id, true); e != nil {
			return e
		}
	}
	return nil
}
