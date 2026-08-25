package service

import (
	"choirsearch/internal/catalog"
	"choirsearch/internal/model"
	"choirsearch/internal/store"
	"context"
)

type Service struct {
	catalog *catalog.Catalog
	store   *store.Store
}

func New(st *store.Store) *Service { return &Service{catalog: catalog.New(st), store: st} }
func (s *Service) RegisterSong(ctx context.Context, x model.Song, ss []model.Score) error {
	return s.catalog.AddSong(ctx, x, ss)
}
func (s *Service) FindScores(ctx context.Context, q model.SearchQuery) ([]model.SearchResult, error) {
	if e := model.ValidateQuery(q); e != nil {
		return nil, e
	}
	q.Voice = model.NormalizeVoice(q.Voice)
	return s.catalog.Search(ctx, q)
}
func (s *Service) Schedule(ctx context.Context, r model.Rehearsal) error {
	return s.store.SaveRehearsal(ctx, r)
}
func (s *Service) AddMember(ctx context.Context, m model.Member) error {
	if !m.IsEligible() {
		return context.Canceled
	}
	return s.store.SaveMember(ctx, m)
}
func (s *Service) Health(ctx context.Context) error { return s.store.Ping(ctx) }
