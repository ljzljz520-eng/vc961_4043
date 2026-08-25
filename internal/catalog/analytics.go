package catalog

import (
	"choirsearch/internal/model"
	"context"
	"sort"
)

type VoiceSummary struct {
	Voice    string
	Songs    int
	Scores   int
	Featured int
}

func (c *Catalog) VoiceSummaries(ctx context.Context) ([]VoiceSummary, error) {
	songs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	m := map[string]*VoiceSummary{}
	for _, s := range songs {
		for _, v := range s.Voices {
			if m[v] == nil {
				m[v] = &VoiceSummary{Voice: v}
			}
			m[v].Songs++
		}
		ss, e := c.st.ScoresForSong(ctx, s.ID)
		if e != nil {
			return nil, e
		}
		for _, z := range ss {
			q := m[z.Voice]
			if q == nil {
				q = &VoiceSummary{Voice: z.Voice}
				m[z.Voice] = q
			}
			q.Scores++
			if z.Featured {
				q.Featured++
			}
		}
	}
	out := []VoiceSummary{}
	for _, x := range m {
		out = append(out, *x)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Voice < out[j].Voice })
	return out, nil
}
func (c *Catalog) DifficultyCounts(ctx context.Context) (map[string]int, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	m := map[string]int{}
	for _, x := range xs {
		m[x.Difficulty]++
	}
	return m, nil
}
func (c *Catalog) FeaturedCount(ctx context.Context) (int, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return 0, e
	}
	n := 0
	for _, x := range xs {
		ss, e := c.st.ScoresForSong(ctx, x.ID)
		if e != nil {
			return n, e
		}
		for _, z := range ss {
			if z.Featured {
				n++
			}
		}
	}
	return n, nil
}
func compareSongs(a, b model.Song) bool { return a.Title < b.Title }
