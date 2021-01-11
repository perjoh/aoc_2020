package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
)

type node struct {
	value int
	prev  *node
	next  *node
	lst   *circularList
}

type circularList struct {
	head   *node
	size   int
	lookup map[int]*node
}

func (lst *circularList) add(value int) {
	newNode := node{value: value}
	newNode.lst = lst
	lst.size++
	if lst.head == nil {
		lst.head = &newNode
		lst.head.prev = &newNode
		lst.head.next = &newNode
		lst.lookup = map[int]*node{}
	} else {
		last := lst.head.prev
		lst.head.prev = &newNode
		last.next = &newNode
		newNode.prev = last
		newNode.next = lst.head
	}

	lst.lookup[value] = &newNode
}

func inRange(begin, end *node, value int) bool {
	for it := begin; it != end; {
		if it.value == value {
			return true
		}
		it = it.next
	}
	return false
}

func (n *node) findValue(value int) *node {
	if fn, found := n.lst.lookup[value]; found {
		return fn
	}
	return nil
}

func (n *node) step(num int) *node {
	for num > 0 {
		n = n.next
		num--
	}
	return n
}

func (n *node) insertAfter(first, last *node) {
	n.next.prev = last
	last.next = n.next
	n.next = first
	first.prev = n
}

func detachRange(first, last *node) {
	first.prev.next = last.next
	last.next.prev = first.prev
	first.prev = nil
	last.next = nil
}

func findTarget(current, pickFirst, pickLast *node) *node {
	for value := current.value - 1; true; value-- {
		if value < 1 {
			value = current.lst.size
		}

		if !inRange(pickFirst, pickLast.next, value) {
			return current.findValue(value)
		}
	}
	return nil
}

func parseCups(input string) []int {
	result := []int{}
	for _, c := range input {
		result = append(result, int(c)-'0')
	}
	return result
}

func crabCups(cups []int, numRounds int) circularList {
	lst := circularList{}
	for _, v := range cups {
		lst.add(v)
	}

	curCup := lst.head
	for numRounds > 0 {

		pickFirst := curCup.next
		pickLast := pickFirst.step(3).prev
		detachRange(pickFirst, pickLast)
		target := findTarget(curCup, pickFirst, pickLast)
		target.insertAfter(pickFirst, pickLast)
		curCup = curCup.next

		numRounds--
	}

	return lst
}

func inList(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func cupLabels(cups circularList) string {
	firstNode := cups.head.findValue(1).next
	lastNode := firstNode.step(cups.size - 1)
	var result string
	for curNode := firstNode; curNode != lastNode; curNode = curNode.next {
		result += strconv.Itoa(curNode.value)
	}
	return result
}

func part1(input []string) string {
	cups := parseCups(input[0])
	lst := crabCups(cups, 100)
	return cupLabels(lst)
}

func appendBigCups(cups []int, size int) []int {
	for i := len(cups) + 1; i <= size; i++ {
		cups = append(cups, i)
	}
	return cups
}

func calcProduct(lst circularList) int {
	n := lst.head.findValue(1)
	return n.next.value * n.next.next.value
}

func part2(input []string) int {
	cups := parseCups(input[0])
	cups = appendBigCups(cups, 1000000)
	lst := crabCups(cups, 10000000)
	return calcProduct(lst)
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
