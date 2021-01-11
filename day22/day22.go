package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

func parseHands(input []string) ([]int, []int) {
	handA := []int{}
	handB := []int{}
	hand := &handA
	for _, line := range input {
		if line == "" {
			hand = &handB
		} else if !strings.Contains(line, "Player") {
			val, _ := strconv.Atoi(line)
			*hand = append(*hand, val)
		}
	}
	return handA, handB
}

func calcScore(hand []int, offset int) int {
	score := 0
	for i, val := range hand {
		score += (len(hand) - i + offset) * val
	}
	return score
}

func playHands(handA []int, handB []int) int {
	for len(handA) > 0 && len(handB) > 0 {
		cardA := handA[0]
		handA = handA[1:]
		cardB := handB[0]
		handB = handB[1:]

		if cardA > cardB {
			handA = append(handA, cardA, cardB)
		} else if cardB > cardA {
			handB = append(handB, cardB, cardA)
		}
	}

	winningHand := &handA
	if len(handA) == 0 {
		winningHand = &handB
	}

	return calcScore(*winningHand, 0)
}

func part1(input []string) int {
	handA, handB := parseHands(input)
	return playHands(handA, handB)
}

func inList(lst []int, val int) bool {
	for _, v := range lst {
		if v == val {
			return true
		}
	}
	return false
}

// Returns: (winner, score)
// 			winner: 0 - draw, 1 - handA, 2 - handB
func playCombat(handA []int, handB []int) (int, int) {

	prevScoreA := []int{}
	prevScoreB := []int{}

	for len(handA) > 0 && len(handB) > 0 {

		// use calcscore as cheap crc
		scoreA := calcScore(handA, 23)
		scoreB := calcScore(handB, 23)

		if inList(prevScoreA, scoreA) && inList(prevScoreB, scoreB) {
			return 1, scoreA
		}

		prevScoreA = append(prevScoreA, scoreA)
		prevScoreB = append(prevScoreB, scoreB)

		cardA := handA[0]
		cardB := handB[0]

		handA = handA[1:]
		handB = handB[1:]

		winner := 0

		if len(handA) >= cardA && len(handB) >= cardB {
			newHandA := append([]int{}, handA[:cardA]...)
			newHandB := append([]int{}, handB[:cardB]...)
			winner, _ = playCombat(newHandA, newHandB)
		} else if cardA > cardB {
			winner = 1
		} else if cardB > cardA {
			winner = 2
		}

		if winner == 1 {
			handA = append(handA, cardA, cardB)
		} else if winner == 2 {
			handB = append(handB, cardB, cardA)
		}
	}

	if len(handA) == 0 {
		return 2, calcScore(handB, 0)
	}
	return 1, calcScore(handA, 0)
}

func part2(input []string) int {
	handA, handB := parseHands(input)
	_, score := playCombat(handA, handB)
	return score
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
