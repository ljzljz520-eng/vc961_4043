package scores

import (
	"choirsearch/internal/model"
	"strings"
)

func Filter(items []model.Score, voice, format string) []model.Score {
	result := items[:0]
	for _, x := range items {
		if voice != "" && !strings.EqualFold(x.Voice, voice) {
			continue
		}
		if format != "" && x.Format != format {
			continue
		}
		result = append(result, x)
	}
	return result
}
func SortByPages(items []model.Score) []model.Score {
	out := append([]model.Score(nil), items...)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j].Pages < out[j-1].Pages; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}
func Featured(items []model.Score) []model.Score {
	out := []model.Score{}
	for _, x := range items {
		if x.Featured {
			out = append(out, x)
		}
	}
	return out
}
func GroupByFormat(items []model.Score) map[string][]model.Score {
	m := map[string][]model.Score{}
	for _, x := range items {
		m[x.Format] = append(m[x.Format], x)
	}
	return m
}
