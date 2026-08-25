package service

import (
	"choirsearch/internal/model"
	"context"
	"sort"
	"time"
)

type Planner struct{ svc *Service }

func NewPlanner(s *Service) *Planner { return &Planner{svc: s} }
func (p *Planner) Agenda(ctx context.Context) ([]model.Rehearsal, error) {
	return p.svc.store.ListRehearsals(ctx)
}
func (p *Planner) Next(ctx context.Context, now time.Time) (model.Rehearsal, error) {
	xs, e := p.Agenda(ctx)
	if e != nil {
		return model.Rehearsal{}, e
	}
	for _, x := range xs {
		if x.StartsAt.After(now) {
			return x, nil
		}
	}
	return model.Rehearsal{}, context.Canceled
}
func (p *Planner) ConfirmRoster(ctx context.Context, r model.Rehearsal) ([]model.Member, error) {
	ms, e := p.svc.store.ActiveMembers(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.Member{}
	for _, m := range ms {
		if m.Voice != "" && r.Conductor != "" {
			out = append(out, m)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}
func (p *Planner) Capacity(ctx context.Context, r model.Rehearsal, limit int) (bool, error) {
	ms, e := p.ConfirmRoster(ctx, r)
	if e != nil {
		return false, e
	}
	return len(ms) <= limit, nil
}
func (p *Planner) Window(r model.Rehearsal) time.Duration { return r.EndAt().Sub(r.StartsAt) }
