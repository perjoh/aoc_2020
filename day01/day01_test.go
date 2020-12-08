package main

import (
	"testing"
)

func TestFindSumEntry(t *testing.T) {
	list := []int{1721, 979, 366, 299, 675, 1456}

	if indexA, indexB, e := findSumEntry(list, 2020); e == nil {
		value := list[indexA] * list[indexB]
		if value != 514579 {
			t.Fatalf("Incorrect value: %d (%d, %d)", value, indexA, indexB)
		}
	} else {
		t.Fatalf("No sum found")
	}
}

func TestFindSumEntryCubic(t *testing.T) {
	list := []int{1721, 979, 366, 299, 675, 1456}

	if indexA, indexB, indexC, e := findSumEntryCubic(list, 2020); e == nil {
		value := list[indexA] * list[indexB] * list[indexC]
		if value != 241861950 {
			t.Fatalf("Incorrect value: %d (%d, %d, %d)", value, indexA, indexB, indexC)
		}
	} else {
		t.Fatalf("No sum found")
	}

}
