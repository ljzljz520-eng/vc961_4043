package scores

import (
	"choirsearch/internal/model"
	"testing"
)

func TestScoreFilterNoStaleItems(t *testing.T) {
	items := []model.Score{{ID: 1, Voice: "S"}, {ID: 2, Voice: "A"}}
	got := Filter(items, "A", "")
	if len(got) != 1 || got[0].ID != 2 {
		t.Fatalf("unexpected %+v", got)
	}
}
