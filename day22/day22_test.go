package main

import "testing"

func TestPlayCombat(t *testing.T) {
	input := []string{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
		"",
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	}

	handA, handB := parseHands(input)
	_, score := playCombat(handA, handB)
	if score != 291 {
		t.Fatalf("playCombat=%v(%v)", score, 291)
	}
}
