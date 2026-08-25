package api

import (
	"choirsearch/internal/model"
	"choirsearch/internal/service"
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct{ s *service.Service }

func New(s *service.Service) *Handler { return &Handler{s: s} }
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	q := model.SearchQuery{Voice: r.URL.Query().Get("voice"), Text: r.URL.Query().Get("q"), Language: r.URL.Query().Get("language"), Difficulty: r.URL.Query().Get("difficulty")}
	x, e := h.s.FindScores(r.Context(), q)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(x)
}
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if e := h.s.Health(r.Context()); e != nil {
		http.Error(w, "down", 503)
		return
	}
	w.WriteHeader(200)
}
func Route(h *Handler) *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/search", h.Search)
	m.HandleFunc("/health", h.Health)
	return m
}
func Request(ctx context.Context, base string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, base+"/health", nil)
	return http.DefaultClient.Do(req)
}
