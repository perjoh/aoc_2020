package main

import (
	"aoc2020/util"
	"fmt"
	"sort"
)

func calcJoltageDiffSum(values []int) (int, int, int) {
	sum := []int{0, 0, 0}
	values = append(values, 0) // add outlet
	sort.Ints(values)
	values = append(values, values[len(values)-1]+3) // add device
	prev := values[0]
	for _, value := range values[1:] {
		if diff := value - prev; diff < 0 || 3 < diff || diff == 0 {
			panic(diff)
		} else {
			sum[diff-1]++
		}
		prev = value
	}

	// charging outlet
	//sum[values[0]-1]++

	// the device
	//sum[3-1]++

	return sum[0], sum[1], sum[2]
}

func partA(input []string) {
	values := util.ToInts(input)
	ones, _, threes := calcJoltageDiffSum(values)
	fmt.Println(ones * threes)
}

// adapters must be sorted
func checkAdapter(adapters []int) int {
	if len(adapters) == 1 {
		return 1
	}
	sum := 0
	joltage := adapters[0]
	tail := adapters[1:]
	for i, value := range tail {
		diff := value - joltage
		if diff > 3 {
			break
		}
		sum += checkAdapter(tail[i:])
	}
	return sum
}

func accumulateAdapters(adapters []int) int64 {
	sort.Ints(adapters)
	sort.Sort(sort.Reverse(sort.IntSlice(adapters)))
	acc := make([]int64, len(adapters))
	acc[0] = 1

	for i := range adapters {
		for j := i + 1; j < len(adapters); j++ {
			diff := adapters[i] - adapters[j] // reverse
			if diff > 3 {
				break
			}
			acc[j] += acc[i]
		}
	}

	return acc[len(acc)-1]
}

func partB(input []string) {
	values := util.ToInts(input)
	values = append(values, 0)
	sort.Ints(values)
	values = append(values, values[len(values)-1]+3)
	//fmt.Println(len(values))
	fmt.Println(accumulateAdapters(values))
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
