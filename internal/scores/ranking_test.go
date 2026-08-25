package scores

import (
	"choirsearch/internal/model"
	"testing"
)

func TestRankingAndPagination(t *testing.T) {
	x := Rank([]model.Score{{ID: 1, Pages: 10}, {ID: 2, Pages: 2, Featured: true}})
	if x[0].ID != 2 {
		t.Fatal(x)
	}
	if len(Paginate(x, 1, 1)) != 1 {
		t.Fatal()
	}
}
