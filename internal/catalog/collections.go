package catalog

import (
	"choirsearch/internal/model"
	"context"
	"sort"
)

func SortSongs(xs []model.Song) []model.Song {
	out := append([]model.Song(nil), xs...)
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return out
}
func FilterDifficulty(xs []model.Song, d string) []model.Song {
	out := []model.Song{}
	for _, x := range xs {
		if x.Difficulty == d {
			out = append(out, x)
		}
	}
	return out
}
func FilterLanguage(xs []model.Song, l string) []model.Song {
	out := []model.Song{}
	for _, x := range xs {
		if x.Language == l {
			out = append(out, x)
		}
	}
	return out
}
func FilterVoice(xs []model.Song, v string) []model.Song {
	out := []model.Song{}
	for _, x := range xs {
		if x.HasVoice(v) {
			out = append(out, x)
		}
	}
	return out
}
func JoinTags(xs []model.Song) map[string]int {
	m := map[string]int{}
	for _, x := range xs {
		for _, t := range x.Tags {
			m[t]++
		}
	}
	return m
}
func PageSongs(xs []model.Song, page, size int) []model.Song {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 25
	}
	a := (page - 1) * size
	if a >= len(xs) {
		return []model.Song{}
	}
	b := a + size
	if b > len(xs) {
		b = len(xs)
	}
	return xs[a:b]
}
func (c *Catalog) Collection(ctx context.Context, name string) (model.Collection, error) {
	xs, e := c.st.ListSongs(ctx)
	if e != nil {
		return model.Collection{}, e
	}
	out := model.Collection{Name: name}
	for _, x := range xs {
		out.Add(x.ID)
	}
	return out, nil
}
func (c *Catalog) CollectionSize(ctx context.Context) (int, error) {
	n, e := c.st.CountSongs(ctx)
	return n, e
}
