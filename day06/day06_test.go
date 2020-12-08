package main

import "testing"

func TestParseGroups(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	groups := parseGroups(input)

	if numGroups := len(groups); numGroups != 5 {
		t.Fatalf("num groups=%d, exp=%d", numGroups, 5)
	}

	expected := 11
	if sum := sumGroups(groups); sum != expected {
		t.Fatalf("act=%d, exp=%d", sum, expected)
	}
}
