package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func EncodeSong(s Song) (string, error) { b, e := json.Marshal(s); return string(b), e }
func DecodeSong(raw string) (Song, error) {
	var s Song
	e := json.Unmarshal([]byte(raw), &s)
	return s, e
}
func EncodeScore(s Score) (string, error) { b, e := json.Marshal(s); return string(b), e }
func DecodeScore(raw string) (Score, error) {
	var s Score
	e := json.Unmarshal([]byte(raw), &s)
	return s, e
}
func ParseSongLine(line string) (Song, error) {
	p := strings.Split(line, "|")
	if len(p) != 4 {
		return Song{}, fmt.Errorf("expected four fields")
	}
	var id int64
	if _, e := fmt.Sscan(p[0], &id); e != nil {
		return Song{}, e
	}
	return NewSong(id, p[1], p[2], strings.Split(p[3], ",")), nil
}
func FormatSongLine(s Song) string {
	return fmt.Sprintf("%d|%s|%s|%s", s.ID, s.Title, s.Composer, strings.Join(s.Voices, ","))
}
func ParseTime(v string) (time.Time, error) { return time.Parse(time.RFC3339, v) }
func FormatTime(v time.Time) string         { return v.UTC().Format(time.RFC3339) }
func CloneSong(s Song) Song {
	out := s
	out.Voices = append([]string(nil), s.Voices...)
	out.Tags = append([]string(nil), s.Tags...)
	return out
}
func CloneScore(s Score) Score { return s }
func EnsureDefaults(s *Song) {
	if s.Language == "" {
		s.Language = "und"
	}
	if s.Difficulty == "" {
		s.Difficulty = "medium"
	}
	s.Tags = NormalizeTags(s.Tags)
}
