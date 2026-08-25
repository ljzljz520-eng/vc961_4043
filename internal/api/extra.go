package api

import (
	"choirsearch/internal/model"
	"choirsearch/internal/service"
	"encoding/json"
	"net/http"
)

func (h *Handler) SearchPost(w http.ResponseWriter, r *http.Request) {
	var q model.SearchQuery
	if e := json.NewDecoder(r.Body).Decode(&q); e != nil {
		WriteJSON(w, 400, ErrorBody{Error: e.Error(), Code: "bad_query"})
		return
	}
	xs, e := h.s.FindScores(r.Context(), q)
	if e != nil {
		WriteJSON(w, 400, ErrorBody{Error: e.Error(), Code: "query_failed"})
		return
	}
	WriteJSON(w, 200, xs)
}
func (h *Handler) Routes() http.Handler                 { return Chain(Route(h)) }
func NewHandlerFromService(s *service.Service) *Handler { return New(s) }
func WriteError(w http.ResponseWriter, status int, code, detail string) {
	WriteJSON(w, status, ErrorBody{Error: detail, Code: code})
}
func DecodeBody(r *http.Request) (model.SearchQuery, error) {
	var q model.SearchQuery
	e := json.NewDecoder(r.Body).Decode(&q)
	return q, e
}
