package main

import "testing"

func TestCountOccupied(t *testing.T) {
	input := []string{
		"#.#L.L#.##",
		"#LLL#LL.L#",
		"L.#.L..#..",
		"#L##.##.L#",
		"#.#L.LL.LL",
		"#.#L#L#.##",
		"..L.L.....",
		"#L#L##L#L#",
		"#.LLLLLL.L",
		"#.#L#L#.##",
	}

	layout := toBytes(input)
	if count := countOccupiedTotal(layout); count != 37 {
		t.Fatalf("countOccupiedTotal %d(%d)", count, 37)
	}
}

func TestProcess(t *testing.T) {
	input := []string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}

	bufA := toBytes(input)
	bufB := make([][]byte, len(bufA))
	for i := range bufB {
		bufB[i] = make([]byte, len(bufA[i]))
		copy(bufB[i], bufA[i])
	}
	//copy(bufB, bufA)

	count := 0
	for process(bufB, bufA, 4, checkOccupied) != 0 {
		bufA, bufB = bufB, bufA
		count++
	}

	if count != 5 {
		t.Fatalf("count %d(%d)", count, 5)
	}

	if count := countOccupiedTotal(bufA); count != 37 {
		t.Fatalf("countOccupiedTotal %d(%d)", count, 37)
	}
}
