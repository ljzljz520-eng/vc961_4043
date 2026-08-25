package store

import (
	"choirsearch/internal/model"
	"context"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/persistent.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveSong(context.Background(), model.NewSong(7, "Reopen", "Composer", []string{"A"})); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	x, e := s.GetSong(context.Background(), 7)
	if e != nil || x.Title != "Reopen" {
		t.Fatalf("%v %+v", e, x)
	}
}
