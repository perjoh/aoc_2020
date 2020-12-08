package main

import (
	"aoc2020/util"
	"fmt"
)

func countTrees(input []string, stepX int, stepY int) int64 {
	var x int
	var numTrees int64
	for y := 0; y < len(input); y += stepY {
		if y > 0 {
			x = (x + stepX) % len(input[y])
			if input[y][x] == '#' {
				numTrees++
			}
		}
	}
	return numTrees
}

func partA(input []string) {
	fmt.Println(countTrees(input, 3, 1))
}

func calcProduct(input []string) int64 {
	return countTrees(input, 1, 1) *
		countTrees(input, 3, 1) *
		countTrees(input, 5, 1) *
		countTrees(input, 7, 1) *
		countTrees(input, 1, 2)
}

func partB(input []string) {
	fmt.Println(calcProduct(input))
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
