package main

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type minMax struct {
	min int
	max int
}

type fieldType struct {
	desc        string
	validRanges []minMax
}

func (field fieldType) validateValue(value int) bool {
	for _, r := range field.validRanges {
		if value <= r.max && r.min <= value {
			return true
		}
	}
	return false
}

func parseFields(input []string) []fieldType {
	result := []fieldType{}
	for _, line := range input {
		if len(line) == 0 {
			break
		}

		r, _ := regexp.Compile("([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
		m := r.FindStringSubmatch(line)
		if m == nil {
			panic("no match")
		}

		min1, _ := strconv.Atoi(m[2])
		max1, _ := strconv.Atoi(m[3])

		min2, _ := strconv.Atoi(m[4])
		max2, _ := strconv.Atoi(m[5])

		result = append(result, fieldType{
			desc:        m[1],
			validRanges: []minMax{{min1, max1}, {min2, max2}}})
	}
	return result
}

func parseTickets(input []string, key string) [][]int {
	result := [][]int{}
	for i := range input {
		if strings.Contains(input[i], key) {
			for _, line := range input[i+1:] {
				if len(line) == 0 {
					break
				}
				result = append(result, util.ToInts(strings.Split(line, ",")))
			}
			break
		}
	}
	return result
}

func validateValue(value int, fields []fieldType) bool {
	for _, field := range fields {
		if field.validateValue(value) {
			return true
		}
	}

	return false
}

func validateTicket(ticket []int, fields []fieldType) int {
	errorRate := 0
	for _, fieldValue := range ticket {
		if !validateValue(fieldValue, fields) {
			errorRate += fieldValue
		}
	}
	return errorRate
}

func validateTickets(tickets [][]int, validRanges []fieldType) int {
	errorRate := 0
	for _, ticket := range tickets {
		errorRate += validateTicket(ticket, validRanges)
	}
	return errorRate
}

func part1(input []string) int {
	fields := parseFields(input)
	nearbyTickets := parseTickets(input, "nearby tickets")
	return validateTickets(nearbyTickets, fields)
}

func removeInvalidTickets(tickets [][]int, fields []fieldType) [][]int {
	for i := 0; i < len(tickets); {
		if validateTicket(tickets[i], fields) != 0 {
			tickets = append(tickets[:i], tickets[i+1:]...) // remove
		} else {
			i++
		}
	}
	return tickets
}

//
func findPossibleFields(tickets [][]int, ticketFields []int, field fieldType) []int {
	result := []int{}
	for _, i := range ticketFields {
		count := 0
		for _, ticket := range tickets {
			if field.validateValue(ticket[i]) {
				count++
			}
		}

		if count == len(tickets) {
			result = append(result, i)
		}
	}
	return result
}

func part2(input []string) int {
	fields := parseFields(input)
	myTicket := parseTickets(input, "your ticket")[0]
	tickets := parseTickets(input, "nearby tickets")
	tickets = append(tickets, myTicket)
	tickets = removeInvalidTickets(tickets, fields)

	determinedFields := make(map[int]fieldType, len(fields))

	ticketFieldsLeft := []int{}
	for i := range tickets[0] {
		ticketFieldsLeft = append(ticketFieldsLeft, i)
	}

	for len(fields) > 0 {
		for i := range fields {
			ticketFields := findPossibleFields(tickets, ticketFieldsLeft, fields[i])
			if len(ticketFields) == 1 {
				determinedFields[ticketFields[0]] = fields[i]
				fields = append(fields[:i], fields[i+1:]...) // Remove it

				for j := range ticketFieldsLeft {
					if ticketFieldsLeft[j] == ticketFields[0] {
						ticketFieldsLeft = append(ticketFieldsLeft[0:j], ticketFieldsLeft[j+1:]...)
						break
					}
				}

				break
			}
		}
	}

	product := 1

	for key, val := range determinedFields {
		if strings.HasPrefix(val.desc, "departure") {
			product *= myTicket[key]
		}
	}

	return product
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
