package main

import (
	"sort"
	"testing"
)

func TestCalcJoltageDiffSum(t *testing.T) {
	values := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	ones, _, threes := calcJoltageDiffSum(values)
	if ones != 7 {
		t.Fatalf("ones=%d, %d", ones, 7)
	}
	if threes != 5 {
		t.Fatalf("threes=%d, %d", threes, 5)
	}
}

func TestCheckAdapters(t *testing.T) {
	values := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
		0,  // outlet
		22, // device
	}

	sort.Ints(values)
	count := checkAdapter(values)
	if count != 8 {
		t.Fatalf("checkAdapter: %d (%d)", count, 8)
	}

	values = []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
		0,  // outlet
		52, // device
	}
	sort.Ints(values)
	count = checkAdapter(values)
	if count != 19208 {
		t.Fatalf("checkAdapter: %d (%d)", count, 19208)
	}
}
