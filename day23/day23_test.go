package main

import "testing"

func TestCrabCups(t *testing.T) {
	cups := parseCups("389125467")
	lst := crabCups(cups, 10)
	expected := "92658374"
	if result := cupLabels(lst); result != expected {
		t.Fatalf("crabCups %v(%v)", result, expected)
	}
}

func TestCrabCupsBig(t *testing.T) {
	cups := parseCups("389125467")
	cups = appendBigCups(cups, 1000000)
	lst := crabCups(cups, 10000000)
	expected := 149245887792
	if result := calcProduct(lst); result != expected {
		t.Fatalf("crabCups %v(%v)", result, expected)
	}
}
