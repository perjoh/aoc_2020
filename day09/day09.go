package main

import (
	"aoc2020/util"
	"fmt"
	"sort"
	"strconv"
)

func toInts(input []string) []int64 {
	result := []int64{}
	for _, val := range input {
		num, e := strconv.ParseInt(val, 10, 64)
		if e != nil {
			panic("too large")
		}
		result = append(result, num)
	}
	return result
}

func findSumVal(values []int64, value int64, sum int64) int {
	for i := range values {
		if values[i]+value == sum {
			return i
		}
	}
	return -1
}

func findSum(values []int64, sum int64) int {
	for i := range values[:len(values)-1] {
		if j := findSumVal(values[i:], values[i], sum); j != -1 {
			return j
		}
	}
	return -1
}

func findInvalid(values []int64, preambleSize int) int {
	for i := preambleSize; i < len(values); i++ {
		if findSum(values[i-preambleSize:i], values[i]) == -1 {
			return i
		}
	}
	return -1
}

func partA(input []string) {
	values := toInts(input)
	fmt.Println(values[findInvalid(values, 25)])
}

type int64slice []int64

func (a int64slice) Len() int           { return len(a) }
func (a int64slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64slice) Less(i, j int) bool { return a[i] < a[j] }

func sum(values []int64) int64 {
	var s int64
	for _, value := range values {
		s += value
	}
	return s
}

func findSequenceSum(values int64slice, valueSum int64) int64 {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if sum(values[i:j]) == valueSum {
				values := append(int64slice{}, values[i:j]...)
				sort.Sort(values)
				return values[0] + values[len(values)-1]
			}
		}
	}
	return 0
}

func partB(input []string) {
	values := toInts(input)
	sum := findSequenceSum(values, values[findInvalid(values, 25)])
	fmt.Println(sum)
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
