package scores

import (
	"choirsearch/internal/model"
	"sort"
)

func Rank(items []model.Score) []model.Score {
	out := append([]model.Score(nil), items...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Featured != out[j].Featured {
			return out[i].Featured
		}
		return out[i].Pages < out[j].Pages
	})
	return out
}
func Paginate(items []model.Score, page, size int) []model.Score {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	start := (page - 1) * size
	if start >= len(items) {
		return []model.Score{}
	}
	end := start + size
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}
func Voices(items []model.Score) []string {
	m := map[string]bool{}
	for _, x := range items {
		m[x.Voice] = true
	}
	out := []string{}
	for v := range m {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
