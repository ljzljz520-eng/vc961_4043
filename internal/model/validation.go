package model

import "fmt"

func ValidateQuery(q SearchQuery) error {
	if q.Limit < 0 {
		return fmt.Errorf("limit cannot be negative")
	}
	if q.Limit == 0 {
		q.Limit = 50
	}
	if len(q.Voice) > 32 {
		return fmt.Errorf("voice too long")
	}
	return nil
}
func NormalizeVoice(v string) string {
	switch v {
	case "soprano", "S":
		return "S"
	case "alto", "A":
		return "A"
	case "tenor", "T":
		return "T"
	case "bass", "B":
		return "B"
	default:
		return v
	}
}
func NormalizeText(s string) string {
	if s == "" {
		return ""
	}
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r += 32
		}
		out = append(out, r)
	}
	return string(out)
}
