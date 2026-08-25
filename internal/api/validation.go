package api

import (
	"choirsearch/internal/model"
	"net/http"
	"strings"
)

func ValidContentType(r *http.Request) bool {
	return strings.Contains(r.Header.Get("content-type"), "application/json")
}
func HeaderOr(r *http.Request, name, value string) string {
	v := r.Header.Get(name)
	if v == "" {
		return value
	}
	return v
}
func ParseVoice(r *http.Request) string    { return model.NormalizeVoice(r.URL.Query().Get("voice")) }
func ParseLanguage(r *http.Request) string { return strings.TrimSpace(r.URL.Query().Get("language")) }
func ParseDifficulty(r *http.Request) string {
	return strings.TrimSpace(r.URL.Query().Get("difficulty"))
}
func ParseText(r *http.Request) string { return strings.TrimSpace(r.URL.Query().Get("q")) }
func QueryFromRequest(r *http.Request) model.SearchQuery {
	return model.SearchQuery{Voice: ParseVoice(r), Language: ParseLanguage(r), Difficulty: ParseDifficulty(r), Text: ParseText(r), Limit: ParseLimit(r)}
}
func WriteNoContent(w http.ResponseWriter) { w.WriteHeader(http.StatusNoContent) }
func IsJSON(r *http.Request) bool {
	return strings.HasSuffix(strings.ToLower(r.URL.Path), ".json") || ValidContentType(r)
}
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		next.ServeHTTP(w, r)
	})
}
func Timeout(next http.Handler) http.Handler { return http.TimeoutHandler(next, 0, "timeout") }
