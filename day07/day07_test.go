package main

import "testing"

func TestParseBag(t *testing.T) {
	bags := parseBags("wavy purple bags contain 1 drab white bag, 4 muted yellow bags, 2 wavy aqua bags.")
	tests := []string{
		"wavy purple",
		"drab white",
		"muted yellow",
		"wavy aqua",
	}

	if numbags := len(bags); numbags != 4 {
		t.Fatalf("numbags, act=%d, exp=%d", numbags, 4)
	}

	for i, test := range tests {
		t.Run(test, func(t *testing.T) {
			if bags[i] != test {
				t.Fatalf("")
			}
		})
	}
}

func TestCountBags(t *testing.T) {
	input := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	bags := buildBagMap(input)

	if count := countBags(bags, "shiny gold", map[string]struct{}{}); count != 4 {
		t.Fatalf("countBags, %d (exp=%d)", count, 4)
	}

	tests := [][]string{
		{"light red", "bright white", "muted yellow"},
		{"dark orange", "bright white", "muted yellow"},
		{"bright white", "shiny gold"},
		{"muted yellow", "shiny gold", "faded blue"},
		{"shiny gold", "dark olive", "vibrant plum"},
		{"dark olive", "faded blue", "dotted black"},
		{"vibrant plum", "faded blue", "dotted black"},
		{"faded blue"},
		{"dotted black"},
	}

	for i, test := range tests {
		t.Run("", func(t *testing.T) {
			bags := parseBags(input[i])
			for j, bag := range bags {
				if len(test) <= j {
					t.Fatalf("index out of range, %s", bag)
				}
				if test[j] != bag {
					t.Fatalf("%s (exp=%s)", bag, tests[i])
				}
			}
		})
	}
}

func TestCountBagsEx(t *testing.T) {
	input := []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}

	tests := []struct {
		parent string
		child  string
		count  int
	}{
		{"shiny gold", "dark red", 2},
		{"dark red", "dark orange", 2},
		{"dark orange", "dark yellow", 2},
		{"dark yellow", "dark green", 2},
		{"dark green", "dark blue", 2},
		{"dark blue", "dark violet", 2},
	}

	tree := buildBagTree(input)

	for _, test := range tests {
		t.Run(test.parent, func(t *testing.T) {
			child := tree[test.parent]
			if !(child[0].desc == test.child && child[0].count == test.count) {
				t.Fatalf("p=%s,ch=%s,c=%d", test.parent, child[0].desc, child[0].count)
			}
		})
	}

	if l := len(tree); l != 7 {
		t.Fatalf("len, %d (exp=%d)", l, 7)
	}

	if count := countBagsEx(tree, "shiny gold") - 1; count != 126 {
		t.Fatalf("%d (exp=%d)", count, 126)
	}

}
