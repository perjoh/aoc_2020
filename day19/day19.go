package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

type messageRule struct {
	char     byte
	rulesets [][]int
}

func parseRuleset(input string) []int {
	rules := strings.Split(strings.TrimSpace(input), " ")
	return util.ToInts(rules)
}

type messageRuleMap map[int]messageRule

func parseRules(input []string) messageRuleMap {
	rules := messageRuleMap{}
	for _, rule := range input {
		if len(rule) == 0 {
			break
		}

		split := strings.Split(rule, ":")
		ruleID, _ := strconv.Atoi(split[0])

		if strings.Contains(split[1], "|") {
			rulesStr := strings.Split(split[1], "|")

			rules[ruleID] = messageRule{
				rulesets: [][]int{parseRuleset(rulesStr[0]), parseRuleset(rulesStr[1])},
			}

		} else if strings.Contains(split[1], "\"") {
			pos := strings.Index(split[1], "\"")
			rules[ruleID] = messageRule{char: byte(split[1][pos+1])}
		} else {
			rules[ruleID] = messageRule{rulesets: [][]int{parseRuleset(split[1])}}
		}
	}
	return rules
}

// Returns if rule passed and "width" of rule in characters.
func (rules messageRuleMap) checkRuleInternal(ruleID int, message string, pos int) (bool, int) {
	if pos < len(message) {
		rule, has := rules[ruleID]
		if has {
			if rule.char != 0 {
				return rule.char == message[pos], 1
			}

			ruleChecker := func(ruleset []int) (bool, int) {
				offset := 0
				for _, subRuleID := range ruleset {
					success, step := rules.checkRuleInternal(subRuleID, message, pos+offset)

					offset += step
					if !success {
						return false, offset
					}
				}
				return true, offset
			}

			for _, ruleset := range rule.rulesets {
				if len(ruleset) > 0 {
					if ok, offset := ruleChecker(ruleset); ok {
						return ok, offset
					}
				}
			}
		}
	}

	return false, 0
}

func (rules messageRuleMap) checkRules(message string) bool {
	if ok, length := rules.checkRuleInternal(0, message, 0); ok && len(message) == length {
		return true
	}

	return false
}

func parseMessages(input []string) []string {
	messages := []string{}
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			for j := i + 1; j < len(input); j++ {
				messages = append(messages, input[j])
			}
		}
	}
	return messages
}

func part1(input []string) int {
	rules := parseRules(input)
	messages := parseMessages(input)

	count := 0
	for _, msg := range messages {
		if rules.checkRules(msg) {
			count++
		}
	}
	return count
}

func part2(input []string) int {
	rules := parseRules(input)

	// try all reasonable combos over messages that did not match

	eightRules := [][]int{
		{42, 42, 42, 42, 42, 42, 42, 42},
		{42, 42, 42, 42, 42, 42, 42},
		{42, 42, 42, 42, 42, 42},
		{42, 42, 42, 42, 42},
		{42, 42, 42, 42},
		{42, 42, 42},
		{42, 42},
		{42},
	}

	elevenRules := [][]int{
		{42, 42, 42, 42, 42, 42, 42, 31, 31, 31, 31, 31, 31, 31},
		{42, 42, 42, 42, 42, 42, 31, 31, 31, 31, 31, 31},
		{42, 42, 42, 42, 42, 31, 31, 31, 31, 31},
		{42, 42, 42, 42, 31, 31, 31, 31},
		{42, 42, 42, 31, 31, 31},
		{42, 42, 31, 31},
		{42, 31},
	}

	messages := parseMessages(input)
	invalidMessages := []string{}

	for _, msg := range messages {
		if !rules.checkRules(msg) {
			invalidMessages = append(invalidMessages, msg)
		}
	}

	newMessages := map[string]struct{}{}
	for _, msg := range invalidMessages {
		for _, eightRule := range eightRules {
			for _, elevenRule := range elevenRules {
				rules[8] = messageRule{rulesets: [][]int{eightRule}}
				rules[11] = messageRule{rulesets: [][]int{elevenRule}}

				if rules.checkRules(msg) {
					newMessages[msg] = struct{}{}
				}
			}
		}
	}

	return len(messages) - len(invalidMessages) + len(newMessages)
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
