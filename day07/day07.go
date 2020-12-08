package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

// wavy purple bags contain 1 drab white bag, 4 muted yellow bags, 2 wavy aqua bags.
func parseBags(input string) []string {
	split := strings.Split(input, " ")
	bags := make([]string, 0)
	bags = append(bags, split[0]+" "+split[1])

	if !strings.Contains(input, "no other bags") {
		for i := 4; i < len(split); i += 4 {
			bags = append(bags, split[i+1]+" "+split[i+2])
		}
	}

	return bags
}

type bagParentMap map[string][]string

func buildBagMap(input []string) bagParentMap {
	bags := bagParentMap{}
	for _, bagData := range input {
		newBags := parseBags(bagData)
		for _, bag := range newBags[1:] {
			bags[bag] = append(bags[bag], newBags[0])
		}
	}
	return bags
}

func countBags(bagMap bagParentMap, bagToCheck string, uniqueBags map[string]struct{}) int {
	bags, _ := bagMap[bagToCheck]
	sum := 0
	for _, bag := range bags {
		if _, has := uniqueBags[bag]; !has {
			uniqueBags[bag] = struct{}{}
			sum += countBags(bagMap, bag, uniqueBags) + 1
		}
	}
	return sum
}

func partA(input []string) {
	bagMap := buildBagMap(input)
	fmt.Println(countBags(bagMap, "shiny gold", make(map[string]struct{})))
}

type bagStruct struct {
	count int
	desc  string
}

func parseBagsEx(input string) []bagStruct {
	split := strings.Split(input, " ")
	bags := []bagStruct{{count: 1, desc: split[0] + " " + split[1]}}
	if !strings.Contains(input, "no other bags") {
		for i := 4; i < len(split); i += 4 {
			count, _ := strconv.Atoi(split[i])
			bags = append(bags, bagStruct{count: count, desc: split[i+1] + " " + split[i+2]})
		}
	}

	return bags
}

type bagTree map[string][]bagStruct

func buildBagTree(input []string) bagTree {
	tree := bagTree{}
	for _, line := range input {
		bags := parseBagsEx(line)
		tree[bags[0].desc] = bags[1:]
	}
	return tree
}

func countBagsEx(bags bagTree, bagDesc string) int {
	sum := 1
	_, has := bags[bagDesc]
	if has {
		for _, bag := range bags[bagDesc] {
			sum += countBagsEx(bags, bag.desc) * bag.count
		}
	}
	return sum
}

func partB(input []string) {
	tree := buildBagTree(input)
	fmt.Println(countBagsEx(tree, "shiny gold") - 1)
}

func main() {
	input := util.ReadStdin()
	partA(input)
	partB(input)
}
