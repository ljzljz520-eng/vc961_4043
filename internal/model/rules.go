package model

import "time"

type Collection struct {
	ID        int64
	Name      string
	Curator   string
	Published bool
	SongIDs   []int64
}
type Annotation struct {
	ID        int64
	SongID    int64
	Author    string
	Text      string
	CreatedAt time.Time
	Resolved  bool
}
type Reservation struct {
	ID          int64
	RehearsalID int64
	MemberID    int64
	Confirmed   bool
}

func (c Collection) Contains(id int64) bool {
	for _, x := range c.SongIDs {
		if x == id {
			return true
		}
	}
	return false
}
func (c *Collection) Add(id int64) bool {
	if c.Contains(id) {
		return false
	}
	c.SongIDs = append(c.SongIDs, id)
	return true
}
func (c *Collection) Remove(id int64) bool {
	for i, x := range c.SongIDs {
		if x == id {
			c.SongIDs = append(c.SongIDs[:i], c.SongIDs[i+1:]...)
			return true
		}
	}
	return false
}
func (a Annotation) IsUsable() bool     { return a.SongID > 0 && a.Author != "" && a.Text != "" }
func (r Reservation) IsConfirmed() bool { return r.Confirmed && r.RehearsalID > 0 && r.MemberID > 0 }
func (s Song) Age(now time.Time) time.Duration {
	if s.CreatedAt.IsZero() {
		return 0
	}
	return now.Sub(s.CreatedAt)
}
func (s Song) SearchBlob() string {
	return s.Title + " " + s.Composer + " " + s.Language + " " + s.Difficulty
}
