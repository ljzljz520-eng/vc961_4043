package service

import (
	"choirsearch/internal/model"
	"context"
	"sort"
	"strings"
)

type Metrics struct {
	Songs        int
	Scores       int
	Voices       map[string]int
	Languages    map[string]int
	Difficulties map[string]int
}

func (s *Service) Metrics(ctx context.Context) (Metrics, error) {
	xs, e := s.store.ListSongs(ctx)
	if e != nil {
		return Metrics{}, e
	}
	m := Metrics{Songs: len(xs), Voices: map[string]int{}, Languages: map[string]int{}, Difficulties: map[string]int{}}
	for _, x := range xs {
		m.Languages[x.Language]++
		m.Difficulties[x.Difficulty]++
		for _, v := range x.Voices {
			m.Voices[v]++
		}
		n, e := s.store.ScoreCount(ctx, x.ID)
		if e != nil {
			return m, e
		}
		m.Scores += n
	}
	return m, nil
}
func (m Metrics) VoiceList() []string {
	out := []string{}
	for v := range m.Voices {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
func (m Metrics) LanguageList() []string {
	out := []string{}
	for v := range m.Languages {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
func (m Metrics) DifficultyList() []string {
	out := []string{}
	for v := range m.Difficulties {
		out = append(out, v)
	}
	sort.Strings(out)
	return out
}
func (m Metrics) HasVoice(v string) bool      { return m.Voices[v] > 0 }
func (m Metrics) HasLanguage(v string) bool   { return m.Languages[v] > 0 }
func (m Metrics) HasDifficulty(v string) bool { return m.Difficulties[v] > 0 }
func (m Metrics) Description() string {
	return strings.Join([]string{"songs", "scores", "voices"}, " ")
}
func (s *Service) VoiceCoverage(ctx context.Context) float64 {
	m, e := s.Metrics(ctx)
	if e != nil {
		return 0
	}
	return float64(len(m.Voices)) / 4
}
func (s *Service) HasCatalog(ctx context.Context) bool {
	m, e := s.Metrics(ctx)
	return e == nil && m.Songs > 0
}
func (s *Service) EnsureSong(ctx context.Context, x model.Song) error {
	if s.HasCatalog(ctx) {
		return nil
	}
	return s.RegisterSong(ctx, x, nil)
}
