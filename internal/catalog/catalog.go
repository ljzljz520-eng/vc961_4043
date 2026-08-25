package catalog

import (
	"choirsearch/internal/model"
	"choirsearch/internal/scores"
	"choirsearch/internal/store"
	"context"
	"strings"
)

type Catalog struct{ st *store.Store }

func New(st *store.Store) *Catalog { return &Catalog{st: st} }
func (c *Catalog) AddSong(ctx context.Context, x model.Song, ss []model.Score) error {
	if !x.IsValid() {
		return context.Canceled
	}
	if e := c.st.SaveSong(ctx, x); e != nil {
		return e
	}
	for _, z := range ss {
		if e := c.st.SaveScore(ctx, z); e != nil {
			return e
		}
	}
	return nil
}
func (c *Catalog) Search(ctx context.Context, q model.SearchQuery) ([]model.SearchResult, error) {
	all, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.SearchResult{}
	needle := strings.ToLower(q.Text)
	for _, x := range all {
		if q.Voice != "" && !x.HasVoice(q.Voice) {
			continue
		}
		if q.Language != "" && x.Language != q.Language {
			continue
		}
		if q.Difficulty != "" && x.Difficulty != q.Difficulty {
			continue
		}
		if needle != "" && !strings.Contains(strings.ToLower(x.Title+" "+x.Composer), needle) {
			continue
		}
		ss, e := c.st.ScoresForSong(ctx, x.ID)
		if e != nil {
			return nil, e
		}
		ss = scores.Filter(ss, q.Voice, "")
		out = append(out, model.SearchResult{Song: x, Scores: ss, MatchReason: "metadata match"})
	}
	if q.Limit > 0 && len(out) > q.Limit {
		out = out[:q.Limit]
	}
	return out, nil
}
func (c *Catalog) Remove(ctx context.Context, id int64) error { return c.st.DeleteSong(ctx, id) }
func (c *Catalog) RehearsalAgenda(ctx context.Context) ([]model.Rehearsal, error) {
	return c.st.ListRehearsals(ctx)
}
