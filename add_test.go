package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 4)
	expected := 6

	if sum != expected {
		t.Errorf("expected %d sum %d", expected, sum)
	}
}
