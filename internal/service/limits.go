package service

import (
	"choirsearch/internal/model"
	"context"
	"fmt"
)

func LimitQuery(q model.SearchQuery, max int) model.SearchQuery {
	if max < 1 {
		max = 1
	}
	if q.Limit == 0 || q.Limit > max {
		q.Limit = max
	}
	return q
}
func RequireVoice(q model.SearchQuery) error {
	if q.Voice == "" {
		return fmt.Errorf("voice required")
	}
	return nil
}
func RequireText(q model.SearchQuery) error {
	if q.Text == "" {
		return fmt.Errorf("text required")
	}
	return nil
}
func (s *Service) SearchLimited(ctx context.Context, q model.SearchQuery, max int) ([]model.SearchResult, error) {
	q = LimitQuery(q, max)
	return s.FindScores(ctx, q)
}
func (s *Service) SearchVoice(ctx context.Context, v string) ([]model.SearchResult, error) {
	return s.FindScores(ctx, model.SearchQuery{Voice: v})
}
func (s *Service) SearchText(ctx context.Context, t string) ([]model.SearchResult, error) {
	return s.FindScores(ctx, model.SearchQuery{Text: t})
}
func (s *Service) SearchLanguage(ctx context.Context, l string) ([]model.SearchResult, error) {
	return s.FindScores(ctx, model.SearchQuery{Language: l})
}
func (s *Service) SearchDifficulty(ctx context.Context, d string) ([]model.SearchResult, error) {
	return s.FindScores(ctx, model.SearchQuery{Difficulty: d})
}
func (s *Service) SearchCombined(ctx context.Context, q model.SearchQuery) ([]model.SearchResult, error) {
	if e := model.ValidateQuery(q); e != nil {
		return nil, e
	}
	return s.FindScores(ctx, q)
}
func (s *Service) EmptySearch(ctx context.Context) (bool, error) {
	r, e := s.FindScores(ctx, model.SearchQuery{Limit: 1})
	return len(r) == 0, e
}
