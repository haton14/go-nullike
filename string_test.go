package nullike

import (
	"testing"
)

func TestValid(t *testing.T) {
	var s NullString
	if s.Valid() {
		t.Fatalf("must be not null have %v", s)
	}
	s = NullString{s: nil}
	if s.Valid() {
		t.Fatalf("must be not null have %v", s)
	}
	tmp := "ohagi"
	s = NullString{s: &tmp}
	if !s.Valid() {
		t.Fatalf("must be not null but have %v", s)
	}
}
