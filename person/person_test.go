package person

import (
	"testing"
)

func TestPersonString(t *testing.T) {
	p := Person{
		Name:   "John",
		Amount: -5.0,
	}
	expected := "John -> -5.00 (Will collect)"

	if p.String() != expected {
		t.Errorf("Expected %s but got %s", expected, p)
	}
}
