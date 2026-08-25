package api

import (
	"choirsearch/internal/service"
	"choirsearch/internal/store"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	s, _ := store.Open(t.TempDir() + "/a.db")
	defer s.Close()
	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	New(service.New(s)).Health(r, req)
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
