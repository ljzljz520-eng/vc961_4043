package catalog

import (
	"choirsearch/internal/model"
	"context"
	"sort"
	"strings"
)

func normalize(s string) string { return strings.ToLower(strings.TrimSpace(s)) }
func includeSong(s model.Song, q model.SearchQuery) bool {
	if q.Voice != "" && !s.HasVoice(q.Voice) {
		return false
	}
	if q.Language != "" && normalize(s.Language) != normalize(q.Language) {
		return false
	}
	if q.Difficulty != "" && normalize(s.Difficulty) != normalize(q.Difficulty) {
		return false
	}
	if q.Text != "" && !strings.Contains(normalize(s.SearchBlob()), normalize(q.Text)) {
		return false
	}
	return true
}
func rankSong(s model.Song, q model.SearchQuery) int {
	score := 0
	if q.Text != "" && normalize(s.Title) == normalize(q.Text) {
		score += 100
	}
	if q.Language != "" {
		score += 10
	}
	if q.Difficulty != "" {
		score += 5
	}
	if s.HasVoice(q.Voice) {
		score += 2
	}
	return score
}
func (c *Catalog) SearchRanked(ctx context.Context, q model.SearchQuery) ([]model.SearchResult, error) {
	all, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.SearchResult{}
	for _, s := range all {
		if !includeSong(s, q) {
			continue
		}
		ss, e := c.st.ScoresForSong(ctx, s.ID)
		if e != nil {
			return nil, e
		}
		out = append(out, model.SearchResult{Song: s, Scores: ss, MatchReason: "ranked"})
	}
	sort.SliceStable(out, func(i, j int) bool { return rankSong(out[i].Song, q) > rankSong(out[j].Song, q) })
	return out, nil
}
func (c *Catalog) Languages(ctx context.Context) ([]string, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	m := map[string]bool{}
	for _, x := range xs {
		if x.Language != "" {
			m[x.Language] = true
		}
	}
	out := []string{}
	for x := range m {
		out = append(out, x)
	}
	sort.Strings(out)
	return out, nil
}
func (c *Catalog) Difficulties(ctx context.Context) ([]string, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	m := map[string]bool{}
	for _, x := range xs {
		if x.Difficulty != "" {
			m[x.Difficulty] = true
		}
	}
	out := []string{}
	for x := range m {
		out = append(out, x)
	}
	sort.Strings(out)
	return out, nil
}
func (c *Catalog) Titles(ctx context.Context, prefix string) ([]string, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	out := []string{}
	for _, x := range xs {
		if strings.HasPrefix(normalize(x.Title), normalize(prefix)) {
			out = append(out, x.Title)
		}
	}
	sort.Strings(out)
	return out, nil
}
func (c *Catalog) ComposerSongs(ctx context.Context, composer string) ([]model.Song, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.Song{}
	for _, x := range xs {
		if strings.EqualFold(x.Composer, composer) {
			out = append(out, x)
		}
	}
	return out, nil
}
func (c *Catalog) VoiceSongs(ctx context.Context, voice string) ([]model.Song, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.Song{}
	for _, x := range xs {
		if x.HasVoice(voice) {
			out = append(out, x)
		}
	}
	return out, nil
}
func (c *Catalog) ScoreCount(ctx context.Context, id int64) (int, error) {
	xs, e := c.st.ScoresForSong(ctx, id)
	return len(xs), e
}
func (c *Catalog) HasScore(ctx context.Context, id int64, voice string) (bool, error) {
	xs, e := c.st.ScoresForSong(ctx, id)
	if e != nil {
		return false, e
	}
	for _, x := range xs {
		if x.Voice == voice {
			return true, nil
		}
	}
	return false, nil
}
func (c *Catalog) FeaturedSongs(ctx context.Context) ([]model.Song, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return nil, e
	}
	out := []model.Song{}
	for _, s := range xs {
		ss, e := c.st.ScoresForSong(ctx, s.ID)
		if e != nil {
			return nil, e
		}
		for _, z := range ss {
			if z.Featured {
				out = append(out, s)
				break
			}
		}
	}
	return out, nil
}
func (c *Catalog) Empty(ctx context.Context) (bool, error) {
	n, e := c.st.CountSongs(ctx)
	return n == 0, e
}
