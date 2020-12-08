package main

import (
	"aoc2020/util"
	"fmt"
	"sort"
)

func parseSeatID(boardingPass string) int {
	sid := 0
	for _, key := range boardingPass {
		sid <<= 1
		if key == 'B' || key == 'R' {
			sid |= 1
		}
	}
	return sid
}

func calcRow(seatID int) int {
	return seatID >> 3
}

func calcColumn(seatID int) int {
	return seatID & 7
}

func partA(input []string) {
	var maxSeatID int
	for _, bp := range input {
		if sid := parseSeatID(bp); maxSeatID < sid {
			maxSeatID = sid
		}
	}
	fmt.Println(maxSeatID)
}

func findMissing(sids []int) int {
	sort.Ints(sids)
	mysid := sids[0]
	for _, sid := range sids {
		if sid != mysid {
			return mysid
		}
		mysid++
	}

	return 0
}

func partB(input []string) {
	sids := make([]int, 0)
	for _, bp := range input {
		sids = append(sids, parseSeatID(bp))
	}

	fmt.Println(findMissing(sids))
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
