package model

import (
	"strings"
	"time"
)

func DisplayTitle(s Song) string {
	if s.Composer == "" {
		return s.Title
	}
	return s.Title + " - " + s.Composer
}
func DisplayVoices(s Song) string { return strings.Join(NormalizeVoices(s.Voices), ", ") }
func DisplayScore(s Score) string { return strings.ToUpper(s.Voice) + " " + strings.ToUpper(s.Format) }
func IsDigital(s Score) bool {
	return strings.EqualFold(s.Format, "pdf") || strings.EqualFold(s.Format, "mxl") || strings.EqualFold(s.Format, "musicxml")
}
func IsPrintable(s Score) bool { return strings.EqualFold(s.Format, "pdf") }
func TagLine(s Song) string    { return strings.Join(NormalizeTags(s.Tags), " #") }
func EmptySong() Song          { return Song{Voices: []string{}} }
func EmptyScore() Score        { return Score{Pages: 0} }
func NewMember(id int64, name, email, voice string) Member {
	return Member{ID: id, Name: name, Email: email, Voice: NormalizeVoice(voice)}
}
func NewRehearsal(id int64, name, venue string, at time.Time, duration int) Rehearsal {
	return Rehearsal{ID: id, Name: name, Venue: venue, StartsAt: at, Duration: duration}
}
