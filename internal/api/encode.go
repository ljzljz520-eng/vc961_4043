package api

import (
	"choirsearch/internal/model"
	"encoding/json"
	"net/http"
)

type ErrorBody struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
func DecodeQuery(r *http.Request) model.SearchQuery {
	return model.SearchQuery{Voice: r.URL.Query().Get("voice"), Text: r.URL.Query().Get("q"), Language: r.URL.Query().Get("language"), Difficulty: r.URL.Query().Get("difficulty")}
}
func ParseLimit(r *http.Request) int {
	v := r.URL.Query().Get("limit")
	if v == "" {
		return 50
	}
	n := 0
	for _, c := range v {
		if c < '0' || c > '9' {
			return 50
		}
		n = n*10 + int(c-'0')
		if n > 1000 {
			return 1000
		}
	}
	return n
}
