package main

import (
	"aoc2020/util"
	"fmt"
)

func transformValue(value, subject int) int {
	return (value * subject) % 20201227
}

func getLoopSize(subject int, target int) int {
	loopSize := 1
	value := 1
	for true {
		value = transformValue(value, subject)
		if value == target {
			break
		}

		loopSize++
	}
	return loopSize
}

func transform(subject int, loopSize int) int {
	value := 1
	for loopSize > 0 {
		value = transformValue(value, subject)
		loopSize--
	}
	return value
}

func part1(input []string) int {
	keys := util.ToInts(input)
	loopSizeA := getLoopSize(7, keys[0])
	return transform(keys[1], loopSizeA)
}

func part2(input []string) int {
	return 0
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
