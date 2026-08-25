package store

import (
	"context"
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

type Store struct {
	db   *sql.DB
	path string
}

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}
	db, e := sql.Open("sqlite", path)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	if e = s.schema(); e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) schema() error {
	_, e := s.db.Exec(`CREATE TABLE IF NOT EXISTS songs(id INTEGER PRIMARY KEY,title TEXT,composer TEXT,language TEXT,difficulty TEXT,voices TEXT,tags TEXT,created_at TEXT);CREATE TABLE IF NOT EXISTS scores(id INTEGER PRIMARY KEY,song_id INTEGER,voice TEXT,format TEXT,uri TEXT,pages INTEGER,featured INTEGER);CREATE TABLE IF NOT EXISTS rehearsals(id INTEGER PRIMARY KEY,name TEXT,venue TEXT,starts_at TEXT,duration INTEGER,conductor TEXT);CREATE TABLE IF NOT EXISTS members(id INTEGER PRIMARY KEY,name TEXT,email TEXT,voice TEXT,joined_at TEXT,active INTEGER);`)
	return e
}
func (s *Store) DB() *sql.DB                    { return s.db }
func (s *Store) Close() error                   { return s.db.Close() }
func (s *Store) Ping(ctx context.Context) error { return s.db.PingContext(ctx) }
func (s *Store) Path() string                   { return s.path }
