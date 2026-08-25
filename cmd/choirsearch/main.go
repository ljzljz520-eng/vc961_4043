package main

import (
	"choirsearch/internal/api"
	"choirsearch/internal/service"
	"choirsearch/internal/store"
	"context"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("CHOIR_DB")
	if path == "" {
		path = "choir.db"
	}
	st, e := store.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	svc := service.New(st)
	log.Println("listening on :8080")
	if e = http.ListenAndServe(":8080", api.Route(api.New(svc))); e != nil && e != http.ErrServerClosed {
		log.Fatal(e)
	}
}
func shutdown(ctx context.Context, st *store.Store) error { return st.Close() }
