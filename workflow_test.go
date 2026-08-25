package choirsearch_test

import (
	"choirsearch/internal/model"
	"choirsearch/internal/service"
	"choirsearch/internal/store"
	"context"
	"strings"
	"testing"
	"time"
)

func TestWorkflowSearchByVoice(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/w.db")
	defer s.Close()
	x := service.New(s)
	x.RegisterSong(context.Background(), model.NewSong(1, "Mass", "Mozart", []string{"S", "A"}), []model.Score{{ID: 1, SongID: 1, Voice: "S"}, {ID: 2, SongID: 1, Voice: "A"}})
	r, e := x.FindScores(context.Background(), model.SearchQuery{Voice: "A"})
	if e != nil || len(r) != 1 || len(r[0].Scores) != 1 || r[0].Scores[0].Voice != "A" {
		t.Fatal(e, r)
	}
}
func TestWorkflowImportAndSearch(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/i.db")
	defer s.Close()
	x := service.New(s)
	if n, e := x.ImportLines(context.Background(), strings.NewReader("2|Anthem|Lee|T\n")); e != nil || n != 1 {
		t.Fatal(e, n)
	}
}
func TestWorkflowScheduleAndMembers(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/r.db")
	defer s.Close()
	x := service.New(s)
	if e := x.Schedule(context.Background(), model.Rehearsal{ID: 1, StartsAt: time.Now(), Duration: 60}); e != nil {
		t.Fatal(e)
	}
	if e := x.AddMember(context.Background(), model.Member{ID: 1, Name: "N", Email: "n@x", Active: true}); e != nil {
		t.Fatal(e)
	}
}
