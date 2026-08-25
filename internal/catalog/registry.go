package catalog

import (
	"choirsearch/internal/model"
	"context"
	"sort"
	"strings"
)

type Registry struct {
	Songs  map[int64]model.Song
	Scores map[int64][]model.Score
}

func NewRegistry() *Registry {
	return &Registry{Songs: map[int64]model.Song{}, Scores: map[int64][]model.Score{}}
}
func (r *Registry) PutSong(s model.Song) { r.Songs[s.ID] = model.CloneSong(s) }
func (r *Registry) PutScores(id int64, xs []model.Score) {
	r.Scores[id] = append([]model.Score(nil), xs...)
}
func (r *Registry) GetSong(id int64) (model.Song, bool) {
	x, ok := r.Songs[id]
	return model.CloneSong(x), ok
}
func (r *Registry) GetScores(id int64) []model.Score {
	return append([]model.Score(nil), r.Scores[id]...)
}
func (r *Registry) Remove(id int64) { delete(r.Songs, id); delete(r.Scores, id) }
func (r *Registry) IDs() []int64 {
	out := []int64{}
	for id := range r.Songs {
		out = append(out, id)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}
func (r *Registry) Search(q model.SearchQuery) []model.Song {
	out := []model.Song{}
	for _, x := range r.Songs {
		if q.Voice != "" && !x.HasVoice(q.Voice) {
			continue
		}
		if q.Text != "" && !strings.Contains(strings.ToLower(x.SearchBlob()), strings.ToLower(q.Text)) {
			continue
		}
		out = append(out, model.CloneSong(x))
	}
	return SortSongs(out)
}
func (r *Registry) Count() int { return len(r.Songs) }
func (r *Registry) ScoreCount() int {
	n := 0
	for _, x := range r.Scores {
		n += len(x)
	}
	return n
}
func (r *Registry) Featured() []model.Score {
	out := []model.Score{}
	for _, xs := range r.Scores {
		for _, x := range xs {
			if x.Featured {
				out = append(out, x)
			}
		}
	}
	return out
}
func (r *Registry) VoiceSet() map[string]bool {
	m := map[string]bool{}
	for _, x := range r.Scores {
		for _, z := range x {
			m[z.Voice] = true
		}
	}
	return m
}
func (r *Registry) Rebuild(xs []model.Song) {
	r.Songs = map[int64]model.Song{}
	for _, x := range xs {
		r.PutSong(x)
	}
}
func (c *Catalog) Snapshot(ctx context.Context) (*Registry, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	r := NewRegistry()
	for _, x := range xs {
		r.PutSong(x)
		ss, e := c.st.ScoresForSong(ctx, x.ID)
		if e != nil {
			return nil, e
		}
		r.PutScores(x.ID, ss)
	}
	return r, nil
}
