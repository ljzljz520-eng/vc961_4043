package scores

import (
	"choirsearch/internal/model"
	"math"
)

func AveragePages(xs []model.Score) float64 {
	if len(xs) == 0 {
		return 0
	}
	return float64(PageTotal(xs)) / float64(len(xs))
}
func VariancePages(xs []model.Score) float64 {
	if len(xs) == 0 {
		return 0
	}
	a := AveragePages(xs)
	v := 0.0
	for _, x := range xs {
		d := float64(x.Pages) - a
		v += d * d
	}
	return v / float64(len(xs))
}
func StddevPages(xs []model.Score) float64 { return math.Sqrt(VariancePages(xs)) }
func Coverage(xs []model.Score, voices []string) float64 {
	if len(voices) == 0 {
		return 1
	}
	m := VoiceCounts(xs)
	n := 0
	for _, v := range voices {
		if m[v] > 0 {
			n++
		}
	}
	return float64(n) / float64(len(voices))
}
func IsBalanced(xs []model.Score) bool {
	m := VoiceCounts(xs)
	if len(m) == 0 {
		return false
	}
	min, max := 0, 0
	for _, n := range m {
		if min == 0 || n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max-min <= 1
}
