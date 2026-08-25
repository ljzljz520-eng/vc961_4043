package model

import (
	"sort"
	"strings"
)

func SortByTitle(xs []Song) []Song {
	out := append([]Song(nil), xs...)
	sort.SliceStable(out, func(i, j int) bool { return strings.ToLower(out[i].Title) < strings.ToLower(out[j].Title) })
	return out
}
func SortByComposer(xs []Song) []Song {
	out := append([]Song(nil), xs...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Composer < out[j].Composer })
	return out
}
func FilterValid(xs []Song) []Song {
	out := []Song{}
	for _, x := range xs {
		if x.IsValid() {
			out = append(out, x)
		}
	}
	return out
}
func FilterPublished(xs []Collection) []Collection {
	out := []Collection{}
	for _, x := range xs {
		if x.Published {
			out = append(out, x)
		}
	}
	return out
}
func VoiceOrder() []string { return []string{"S", "A", "T", "B"} }
func VoiceIndex(v string) int {
	for i, x := range VoiceOrder() {
		if x == v {
			return i
		}
	}
	return -1
}
func IsStandardVoice(v string) bool { return VoiceIndex(v) >= 0 }
func NormalizeVoices(xs []string) []string {
	out := []string{}
	for _, x := range xs {
		v := NormalizeVoice(x)
		if IsStandardVoice(v) {
			out = append(out, v)
		}
	}
	return UniqueStrings(out)
}
