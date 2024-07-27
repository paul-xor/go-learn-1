package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(3, 3)
	if total != 6 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.\n", total, 6)
	}
}

func TestSubtract(t *testing.T) {
	result := subtract(15, 10)
	if result != 5 {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.\n", result, 5)
	}
}
