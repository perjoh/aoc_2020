package main

import (
	"aoc2020/util"
	"errors"
	"fmt"
	"sort"
)

func findSumEntry(list []int, desiredSum int) (int, int, error) {
	for i, valueA := range list[:len(list)-1] {
		for j, valueB := range list[i+1:] {
			if valueA+valueB == desiredSum {
				return i, j + i + 1, nil
			}
		}
	}

	return -1, -1, errors.New("no two values found with desired sum")
}

func partA(input []int) {
	if indexA, indexB, e := findSumEntry(input, 2020); e == nil {
		product := input[indexA] * input[indexB]
		fmt.Println(product)
	} else {
		fmt.Println("Error: ", e)
	}
}

func findSumEntryCubic(list []int, desiredSum int) (int, int, int, error) {
	for i, valueA := range list[:len(list)-2] {
		for j, valueB := range list[i+1 : len(list)-1] {
			j = j + i + 1
			for k, valueC := range list[j+1:] {
				k = k + j + 1
				sum := valueA + valueB + valueC
				if sum == desiredSum {
					return i, j, k, nil
				}
			}
		}
	}

	return -1, -1, -1, errors.New("no three values found with desired sum")

}

func partB(input []int) {
	sort.Ints(input)
	i, j, k, e := findSumEntryCubic(input, 2020)
	if e == nil {
		product := input[i] * input[j] * input[k]
		fmt.Println(product)
	} else {
		fmt.Println("Error: ", e)
	}
}

func main() {
	input := util.ToInts(util.ReadStdin())
	partA(input)
	partB(input)
}
