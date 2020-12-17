package main

import (
	"aoc2020/util"
	"fmt"
)

func checkOccupied(layout [][]byte, x int, y int, stepx int, stepy int) bool {
	x += stepx
	y += stepy
	if y < len(layout) && 0 <= y {
		if x < len(layout[0]) && 0 <= x {
			return layout[y][x] == '#'
		}
	}
	return false
}

type occupiedFn func([][]byte, int, int, int, int) bool

func countOccupiedSurrounding(layout [][]byte, x int, y int, occupied occupiedFn) int {
	adjacent := []struct {
		x int
		y int
	}{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 1},
		{0, 1},
		{1, 1},
		{-1, 0},
		{1, 0},
	}

	result := 0
	for _, pt := range adjacent {
		if occupied(layout, x, y, pt.x, pt.y) {
			result++
		}
	}
	return result
}

func toBytes(strings []string) [][]byte {
	buf := make([][]byte, len(strings))
	for i, s := range strings {
		buf[i] = []byte(s)
	}
	return buf
}

// Returns the number of changes.
func process(dst [][]byte, src [][]byte, tolerance int, occupied occupiedFn) int {
	changes := 0
	for y := range src {
		for x := range src[y] {
			if src[y][x] == 'L' {
				if countOccupiedSurrounding(src, x, y, occupied) == 0 {
					dst[y][x] = '#'
					changes++
				} else {
					dst[y][x] = 'L'
				}
			} else if src[y][x] == '#' {
				if countOccupiedSurrounding(src, x, y, occupied) >= tolerance {
					dst[y][x] = 'L'
					changes++
				} else {
					dst[y][x] = '#'
				}
			}
		}
	}
	return changes
}

func countOccupiedTotal(layout [][]byte) int {
	count := 0
	for _, col := range layout {
		for _, c := range col {
			if c == '#' {
				count++
			}
		}
	}
	return count
}

func cloneBuf(buf [][]byte) [][]byte {
	clone := make([][]byte, len(buf))
	for i := range buf {
		clone[i] = make([]byte, len(buf[i]))
		copy(clone[i], buf[i])
	}
	return clone
}

func part1(input []string) int {
	bufA := toBytes(input)
	bufB := cloneBuf(bufA)

	for process(bufB, bufA, 4, checkOccupied) != 0 {
		bufA, bufB = bufB, bufA
	}

	return countOccupiedTotal(bufA)
}

func checkOccupiedLos(layout [][]byte, x int, y int, stepx int, stepy int) bool {
	x += stepx
	y += stepy
	if y < len(layout) && 0 <= y {
		if x < len(layout[0]) && 0 <= x {
			c := layout[y][x]
			if c == '#' {
				return true
			} else if c == 'L' {
				return false
			} else {
				return checkOccupiedLos(layout, x, y, stepx, stepy)
			}
		}
	}
	return false
}

func part2(input []string) int {
	bufA := toBytes(input)
	bufB := cloneBuf(bufA)

	for process(bufB, bufA, 5, checkOccupiedLos) != 0 {
		bufA, bufB = bufB, bufA
	}

	return countOccupiedTotal(bufA)
}

func main() {
	input := util.ReadStdin()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
