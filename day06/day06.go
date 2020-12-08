package main

import (
	"aoc2020/util"
	"fmt"
	"strings"
)

type group struct {
	size    int
	answers map[rune]int
}

func parseGroup(answers []string) group {
	result := group{len(answers), make(map[rune]int)}
	for _, answer := range answers {
		for _, yes := range answer {
			val, _ := result.answers[yes]
			result.answers[yes] = val + 1
		}
	}
	return result
}

func parseGroups(input []string) []group {
	groups := make([]group, 0)
	begin := 0
	for end, line := range input {
		if strings.TrimSpace(line) == "" {
			groups = append(groups, parseGroup(input[begin:end]))
			begin = end + 1
		}
	}

	groups = append(groups, parseGroup(input[begin:]))

	return groups
}

func sumGroups(groups []group) int {
	sum := 0
	for _, group := range groups {
		sum += len(group.answers)
	}
	return sum
}

func partA(input []string) {
	groups := parseGroups(input)
	fmt.Println(sumGroups(groups))
}

func countAllYes(groups []group) int {
	count := 0
	for _, group := range groups {
		for _, val := range group.answers {
			if val == group.size {
				count++
			}
		}
	}
	return count
}

func partB(input []string) {
	groups := parseGroups(input)
	fmt.Println(countAllYes(groups))
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
