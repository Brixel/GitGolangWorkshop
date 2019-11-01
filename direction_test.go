package ggworkshop

import "testing"

func TestContains(t *testing.T) {
	ds := Directions{
		North,
		East,
	}

	if ds.Contains(South) {
		t.Errorf("Directions does not contain South")
	}

	if !ds.Contains(North) {
		t.Errorf("Directions does contain North")
	}
}
