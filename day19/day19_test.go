package main

import "testing"

func TestCheckRule(t *testing.T) {
	input := []string{
		"0: 4 1 5",
		"1: 2 3 | 3 2",
		"2: 4 4 | 5 5",
		"3: 4 5 | 5 4",
		"4: \"a\"",
		"5: \"b\"",
		"",
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}

	rules := parseRules(input)
	messages := parseMessages(input)

	count := 0
	for _, msg := range messages {
		if rules.checkRules(msg) {
			count++
		}
	}

	if count != 2 {
		t.Fatalf("%v (%v)", count, 2)
	}
}
