package catalog

import (
	"context"
	"fmt"
	"sort"
)

type Report struct {
	TotalSongs   int
	TotalScores  int
	ByVoice      map[string]int
	Difficulties map[string]int
}

func (c *Catalog) BuildReport(ctx context.Context) (Report, error) {
	songs, e := c.st.ListSongs(ctx)
	if e != nil {
		return Report{}, e
	}
	r := Report{TotalSongs: len(songs), ByVoice: map[string]int{}, Difficulties: map[string]int{}}
	for _, x := range songs {
		r.Difficulties[x.Difficulty]++
		ss, e := c.st.ScoresForSong(ctx, x.ID)
		if e != nil {
			return r, e
		}
		r.TotalScores += len(ss)
		for _, v := range x.Voices {
			r.ByVoice[v]++
		}
	}
	return r, nil
}
func (r Report) VoiceRanking() []string {
	out := []string{}
	for v := range r.ByVoice {
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return r.ByVoice[out[i]] > r.ByVoice[out[j]] })
	return out
}
func (r Report) Summary() string {
	return fmt.Sprintf("%d songs, %d scores", r.TotalSongs, r.TotalScores)
}
