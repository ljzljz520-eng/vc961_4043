package service

import (
	"choirsearch/internal/store"
	"context"
	"strings"
	"testing"
)

func TestServiceImport(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/s.db")
	defer s.Close()
	x := New(s)
	n, e := x.ImportLines(context.Background(), strings.NewReader("1|Song|Comp|A\n"))
	if e != nil || n != 1 {
		t.Fatal(e, n)
	}
}
