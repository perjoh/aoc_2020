package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input      string
		lastSpoken int
	}{
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if result := part1([]string{test.input}); result != test.lastSpoken {
				t.Fatalf("%s %d (%d)", test.input, result, test.lastSpoken)
			}
		})
	}
}
