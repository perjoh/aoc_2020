package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

func waitTime(busID int, timestamp int) int {
	return busID - timestamp%busID
}

func part1(input []string) int {
	timestamp, _ := strconv.Atoi(input[0])
	buses := strings.Split(input[1], ",")
	earliestBusID := 0

	for _, bus := range buses {
		if bus != "x" {
			busID, _ := strconv.Atoi(bus)
			if earliestBusID == 0 || waitTime(busID, timestamp) < waitTime(earliestBusID, timestamp) {
				earliestBusID = busID
			}
		}
	}
	return earliestBusID * waitTime(earliestBusID, timestamp)
}

type busType struct {
	busID int64
	index int64
}

func parseInput(input []string) []busType {
	buses := []busType{}
	busesInput := strings.Split(input[1], ",")
	for i, bus := range busesInput {
		if bus != "x" {
			busID, _ := strconv.Atoi(bus)
			buses = append(buses, busType{busID: int64(busID), index: int64(i)})
		}
	}
	return buses
}

func findValue(input []string) int64 {
	buses := parseInput(input)
	val := buses[0].busID
	step := val
	for _, bus := range buses[1:] {
		for (val+bus.index)%bus.busID != 0 {
			val += step
		}
		step *= bus.busID
	}
	return val
}

func part2(input []string) int64 {
	return findValue(input)
}

func main() {
	input := util.ReadStdin()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
