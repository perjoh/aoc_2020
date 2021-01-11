package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr   string
		result int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			//(expr := strings.ReplaceAll(test.expr, " ", "")
			if result := eval(test.expr); result != test.result {
				t.Fatalf("%v=%v(%d)", test.expr, result, test.result)
			}
		})
	}
}

func TestEvalFunky(t *testing.T) {
	tests := []struct {
		expr   string
		result int
	}{
		{"2 * (4*5) + 3", 46},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			if result := evalFunky(test.expr); result != test.result {
				t.Fatalf("%v: %v(%v)", test.expr, result, test.result)
			}
		})
	}
}
