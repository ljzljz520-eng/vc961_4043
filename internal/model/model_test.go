package model

import "testing"

func TestModelNormalization(t *testing.T) {
	if NormalizeVoice("soprano") != "S" {
		t.Fatal()
	}
	if NormalizeText("ABC") != "abc" {
		t.Fatal()
	}
	if DifficultyRank("hard") != 3 {
		t.Fatal()
	}
}
