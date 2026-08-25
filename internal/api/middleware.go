package api

import (
	"log"
	"net/http"
	"time"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("x-request-id")
		if id == "" {
			id = time.Now().UTC().Format("20060102150405.000")
		}
		w.Header().Set("x-request-id", id)
		next.ServeHTTP(w, r)
	})
}
func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
func Method(allowed string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowed {
			w.Header().Set("allow", allowed)
			http.Error(w, "method not allowed", 405)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func Chain(h http.Handler) http.Handler { return AccessLog(RequestID(h)) }
