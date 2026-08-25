package model

import (
	"sort"
	"strings"
)

type Facet struct {
	Name   string
	Values []string
}

func QueryKey(q SearchQuery) string {
	return strings.Join([]string{q.Voice, q.Language, q.Difficulty, strings.ToLower(q.Text)}, "|")
}
func QueryEqual(a, b SearchQuery) bool                   { return QueryKey(a) == QueryKey(b) && a.Limit == b.Limit }
func QueryWithVoice(q SearchQuery, v string) SearchQuery { q.Voice = NormalizeVoice(v); return q }
func QueryWithText(q SearchQuery, t string) SearchQuery  { q.Text = strings.TrimSpace(t); return q }
func QueryWithLimit(q SearchQuery, n int) SearchQuery {
	if n < 0 {
		n = 0
	}
	q.Limit = n
	return q
}
func Facets(songs []Song) []Facet {
	langs := map[string]bool{}
	diffs := map[string]bool{}
	voices := map[string]bool{}
	for _, x := range songs {
		langs[x.Language] = true
		diffs[x.Difficulty] = true
		for _, v := range x.Voices {
			voices[v] = true
		}
	}
	return []Facet{{Name: "language", Values: keys(langs)}, {Name: "difficulty", Values: keys(diffs)}, {Name: "voice", Values: keys(voices)}}
}
func keys(m map[string]bool) []string {
	out := []string{}
	for x := range m {
		if x != "" {
			out = append(out, x)
		}
	}
	sort.Strings(out)
	return out
}
func UniqueStrings(xs []string) []string {
	m := map[string]bool{}
	for _, x := range xs {
		m[x] = true
	}
	return keys(m)
}
func ContainsAny(s string, terms []string) bool {
	for _, x := range terms {
		if strings.Contains(strings.ToLower(s), strings.ToLower(x)) {
			return true
		}
	}
	return false
}
func ContainsAll(s string, terms []string) bool {
	for _, x := range terms {
		if !strings.Contains(strings.ToLower(s), strings.ToLower(x)) {
			return false
		}
	}
	return true
}
