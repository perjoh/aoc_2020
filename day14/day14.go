package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseBinaryStr(input string) uint {
	var val uint = 0
	for _, c := range input {
		val <<= 1
		if c == '1' {
			val |= 1
		}
	}
	return val
}

func parseMask(input string, replaceX string) uint {
	r, _ := regexp.Compile("mask = ([01X]+)")
	m := r.FindStringSubmatch(input)
	binaryStr := strings.ReplaceAll(m[1], "X", replaceX)
	return parseBinaryStr(binaryStr)
}

func parseMem(input string) (uint, uint) {
	r, _ := regexp.Compile("mem\\[([0-9]+)\\] = ([0-9]+)")
	m := r.FindStringSubmatch(input)
	address, _ := strconv.Atoi(m[1])
	value, _ := strconv.Atoi(m[2])
	return uint(address), uint(value)
}

func part1(input []string) uint {
	mem := make(map[uint]uint)
	var andMask uint = 0
	var orMask uint = 0
	for _, line := range input {
		if line[1] == 'a' {
			andMask = parseMask(line, "1")
			orMask = parseMask(line, "0")
		} else if line[1] == 'e' {
			address, value := parseMem(line)
			value = (value & andMask) | orMask
			mem[address] = value
		}
	}

	var sum uint = 0
	for _, val := range mem {
		sum += val
	}
	return sum
}

func part2(input []string) uint {
	mem := make(map[uint]uint)
	var mask uint = 0
	var countMask uint = 0
	for _, line := range input {
		if line[1] == 'a' {
			countMaskStr := strings.ReplaceAll(line, "1", "0")
			countMaskStr = strings.ReplaceAll(countMaskStr, "X", "1")
			countMask = ^parseBinaryStr(countMaskStr)
			mask = parseMask(line, "0")
		} else if line[1] == 'e' {
			address, value := parseMem(line)
			count := countMask
			for true {
				newAddress := ((address | mask) & countMask) | (count &^ countMask)
				mem[newAddress] = value
				count = (count + 1) | countMask
				if count == countMask {
					break
				}
			}
		}
	}

	var sum uint = 0
	for _, val := range mem {
		sum += val
	}
	return sum
}

func main() {
	input := util.ReadStdin()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
