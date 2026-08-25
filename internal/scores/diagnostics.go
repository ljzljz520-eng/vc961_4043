package scores

import (
	"choirsearch/internal/model"
	"fmt"
	"sort"
)

func Validate(items []model.Score) []error {
	out := []error{}
	for _, x := range items {
		if x.ID <= 0 {
			out = append(out, fmt.Errorf("missing id"))
		}
		if x.SongID <= 0 {
			out = append(out, fmt.Errorf("missing song"))
		}
		if x.Voice == "" {
			out = append(out, fmt.Errorf("missing voice"))
		}
		if x.Pages < 0 {
			out = append(out, fmt.Errorf("negative pages"))
		}
	}
	return out
}
func VoiceCounts(items []model.Score) map[string]int {
	m := map[string]int{}
	for _, x := range items {
		m[x.Voice]++
	}
	return m
}
func FormatCounts(items []model.Score) map[string]int {
	m := map[string]int{}
	for _, x := range items {
		m[x.Format]++
	}
	return m
}
func PageTotal(items []model.Score) int {
	n := 0
	for _, x := range items {
		n += x.Pages
	}
	return n
}
func MaxPages(items []model.Score) int {
	n := 0
	for _, x := range items {
		if x.Pages > n {
			n = x.Pages
		}
	}
	return n
}
func MinPages(items []model.Score) int {
	if len(items) == 0 {
		return 0
	}
	n := items[0].Pages
	for _, x := range items[1:] {
		if x.Pages < n {
			n = x.Pages
		}
	}
	return n
}
func SortedVoices(items []model.Score) []string {
	m := VoiceCounts(items)
	out := []string{}
	for v := range m {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
