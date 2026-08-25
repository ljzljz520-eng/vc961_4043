package model

import "time"

type Song struct {
	ID                                    int64
	Title, Composer, Language, Difficulty string
	Voices                                []string
	Tags                                  []string
	CreatedAt                             time.Time
}
type Score struct {
	ID                 int64
	SongID             int64
	Voice, Format, URI string
	Pages              int
	Featured           bool
}
type Rehearsal struct {
	ID          int64
	Name, Venue string
	StartsAt    time.Time
	Duration    int
	Conductor   string
}
type Member struct {
	ID                 int64
	Name, Email, Voice string
	JoinedAt           time.Time
	Active             bool
}
type SearchQuery struct {
	Voice, Language, Difficulty, Text string
	Limit                             int
}
type SearchResult struct {
	Song        Song
	Scores      []Score
	MatchReason string
}

func NewSong(id int64, title, composer string, voices []string) Song {
	return Song{ID: id, Title: title, Composer: composer, Voices: voices, CreatedAt: time.Now()}
}
func (s Song) HasVoice(v string) bool {
	for _, x := range s.Voices {
		if x == v {
			return true
		}
	}
	return false
}
func (s Song) IsValid() bool         { return s.ID > 0 && s.Title != "" && len(s.Voices) > 0 }
func (m Member) IsEligible() bool    { return m.Active && m.Name != "" && m.Email != "" }
func (r Rehearsal) EndAt() time.Time { return r.StartsAt.Add(time.Duration(r.Duration) * time.Minute) }
