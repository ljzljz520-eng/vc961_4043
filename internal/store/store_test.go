package store

import (
	"choirsearch/internal/model"
	"context"
	"os"
	"testing"
	"time"
)

func TestStoreSongRoundTrip(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	x := model.NewSong(1, "Missa", "Bach", []string{"S"})
	if e = s.SaveSong(context.Background(), x); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetSong(context.Background(), 1); e != nil {
		t.Fatal(e)
	}
}
func TestStoreOperations(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, _ := Open(p)
	defer s.Close()
	if e := s.SaveRehearsal(context.Background(), model.Rehearsal{ID: 1, StartsAt: time.Now()}); e != nil {
		t.Fatal(e)
	}
	if e := s.SaveMember(context.Background(), model.Member{ID: 1, Name: "A", Email: "a", Active: true}); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Stats(context.Background()); e != nil {
		t.Fatal(e)
	}
	_ = os.Remove(p)
}
