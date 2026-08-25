package service

import (
	"choirsearch/internal/model"
	"context"
	"fmt"
)

func CheckSong(s model.Song) error {
	if s.ID <= 0 {
		return fmt.Errorf("id")
	}
	if s.Title == "" {
		return fmt.Errorf("title")
	}
	if len(s.Voices) == 0 {
		return fmt.Errorf("voices")
	}
	return nil
}
func CheckMember(m model.Member) error {
	if !m.IsEligible() {
		return fmt.Errorf("member")
	}
	return nil
}
func CheckRehearsal(r model.Rehearsal) error {
	if r.ID <= 0 {
		return fmt.Errorf("id")
	}
	if r.Duration <= 0 {
		return fmt.Errorf("duration")
	}
	return nil
}
func CheckScore(s model.Score) error {
	if s.ID <= 0 || s.SongID <= 0 {
		return fmt.Errorf("reference")
	}
	if s.Voice == "" {
		return fmt.Errorf("voice")
	}
	return nil
}
func (s *Service) ValidateCatalog(ctx context.Context) (int, error) {
	xs, e := s.store.ListSongs(ctx)
	if e != nil {
		return 0, e
	}
	n := 0
	for _, x := range xs {
		if CheckSong(x) == nil {
			n++
		}
	}
	return n, nil
}
func (s *Service) ValidateMembers(ctx context.Context) (int, error) {
	xs, e := s.store.ActiveMembers(ctx)
	if e != nil {
		return 0, e
	}
	n := 0
	for _, x := range xs {
		if CheckMember(x) == nil {
			n++
		}
	}
	return n, nil
}
func (s *Service) ValidateRehearsals(ctx context.Context) (int, error) {
	xs, e := s.store.ListRehearsals(ctx)
	if e != nil {
		return 0, e
	}
	n := 0
	for _, x := range xs {
		if CheckRehearsal(x) == nil {
			n++
		}
	}
	return n, nil
}
