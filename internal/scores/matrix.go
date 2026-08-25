package scores

import (
	"choirsearch/internal/model"
	"sort"
	"strings"
)

type Matrix struct{ Rows map[string][]model.Score }

func BuildMatrix(items []model.Score) Matrix {
	m := Matrix{Rows: map[string][]model.Score{}}
	for _, x := range items {
		m.Rows[x.Voice] = append(m.Rows[x.Voice], x)
	}
	return m
}
func (m Matrix) Voices() []string {
	out := []string{}
	for v := range m.Rows {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
func (m Matrix) ForVoice(v string) []model.Score {
	if x, ok := m.Rows[v]; ok {
		return append([]model.Score(nil), x...)
	}
	return []model.Score{}
}
func (m Matrix) Count(v string) int { return len(m.Rows[v]) }
func Merge(a, b []model.Score) []model.Score {
	out := append([]model.Score(nil), a...)
	seen := map[int64]bool{}
	for _, x := range out {
		seen[x.ID] = true
	}
	for _, x := range b {
		if !seen[x.ID] {
			out = append(out, x)
			seen[x.ID] = true
		}
	}
	return out
}
func Deduplicate(items []model.Score) []model.Score {
	out := []model.Score{}
	seen := map[int64]bool{}
	for _, x := range items {
		if !seen[x.ID] {
			out = append(out, x)
			seen[x.ID] = true
		}
	}
	return out
}
func MatchFormat(items []model.Score, f string) []model.Score {
	out := []model.Score{}
	for _, x := range items {
		if strings.EqualFold(x.Format, f) {
			out = append(out, x)
		}
	}
	return out
}
