package main

import (
	"aoc2020/util"
	"fmt"
	"strings"
)

type numberData struct {
	lastTurn int
	prevTurn int
}

func lastNumberSpoken(inputNumbers []int, numTurns int) int {
	numbers := make(map[int]numberData)
	for i, number := range inputNumbers {
		numbers[number] = numberData{lastTurn: i + 1}
	}

	lastSpoken := inputNumbers[len(inputNumbers)-1]

	for turn := len(inputNumbers) + 1; turn <= numTurns; turn++ {
		data, _ := numbers[lastSpoken]
		if data.prevTurn == 0 { // First time spoken
			lastSpoken = 0
		} else {
			lastSpoken = data.lastTurn - data.prevTurn
		}

		data, contains := numbers[lastSpoken]
		if !contains {
			numbers[lastSpoken] = numberData{lastTurn: turn}

		} else {
			data, _ = numbers[lastSpoken]
			numbers[lastSpoken] = numberData{lastTurn: turn, prevTurn: data.lastTurn}
		}
	}
	return lastSpoken
}

func part1(input []string) int {
	input = strings.Split(input[0], ",")
	inputNumbers := util.ToInts(input)
	return lastNumberSpoken(inputNumbers, 2020)
}

func part2(input []string) int {
	input = strings.Split(input[0], ",")
	inputNumbers := util.ToInts(input)
	return lastNumberSpoken(inputNumbers, 30000000)
}

func main() {
	input := util.ReadStdin()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
