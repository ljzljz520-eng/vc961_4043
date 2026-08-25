package service

import (
	"bufio"
	"choirsearch/internal/model"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func (s *Service) ImportLines(ctx context.Context, rd io.Reader) (int, error) {
	sc := bufio.NewScanner(rd)
	n := 0
	for sc.Scan() {
		parts := strings.Split(sc.Text(), "|")
		if len(parts) < 4 {
			return n, fmt.Errorf("line %d malformed", n+1)
		}
		id, e := strconv.ParseInt(parts[0], 10, 64)
		if e != nil {
			return n, e
		}
		x := model.NewSong(id, parts[1], parts[2], strings.Split(parts[3], ","))
		if e = s.RegisterSong(ctx, x, nil); e != nil {
			return n, e
		}
		n++
	}
	return n, sc.Err()
}
func (s *Service) ExportSong(ctx context.Context, id int64) (string, error) {
	x, e := s.store.GetSong(ctx, id)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("%d|%s|%s|%s", x.ID, x.Title, x.Composer, strings.Join(x.Voices, ",")), nil
}
