package model

import (
	"sort"
	"strings"
)

func NormalizeTags(tags []string) []string {
	m := map[string]bool{}
	for _, t := range tags {
		t = strings.TrimSpace(strings.ToLower(t))
		if t != "" {
			m[t] = true
		}
	}
	out := []string{}
	for t := range m {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}
func ScoreLabel(s Score) string {
	if s.Featured {
		return "featured " + s.Format
	}
	return s.Format
}
func DifficultyRank(v string) int {
	switch strings.ToLower(v) {
	case "easy":
		return 1
	case "medium":
		return 2
	case "hard":
		return 3
	default:
		return 0
	}
}
