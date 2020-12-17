package main

import "testing"

func TestPart2(t *testing.T) {
	input := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	if result := part2(input); result != 208 {
		t.Fatalf("part2 %d(%d)", result, 208)
	}
}
