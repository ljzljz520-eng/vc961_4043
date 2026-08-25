package catalog

import (
	"choirsearch/internal/model"
	"choirsearch/internal/store"
	"context"
	"testing"
)

func TestCatalogSearch(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/c.db")
	defer s.Close()
	c := New(s)
	c.AddSong(context.Background(), model.NewSong(1, "Gloria", "Vivaldi", []string{"S"}), []model.Score{{ID: 1, SongID: 1, Voice: "S"}})
	r, e := c.Search(context.Background(), model.SearchQuery{Voice: "S"})
	if e != nil || len(r) != 1 {
		t.Fatal(e, r)
	}
}
