package main

import (
	"fmt"
	"testing"
)

func TestCountTrees(t *testing.T) {
	input := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#"}

	tests := []struct {
		stepX    int
		stepY    int
		expected int64
	}{{1, 1, 2}, {3, 1, 7}, {5, 1, 3}, {7, 1, 4}, {1, 2, 2}}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			if actual := countTrees(input, test.stepX, test.stepY); actual != test.expected {
				t.Fatalf("countTrees, expected: %d, actual: %d", test.expected, actual)
			}
		})
	}
}
