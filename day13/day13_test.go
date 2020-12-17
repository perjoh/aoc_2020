package main

import "testing"

func TestFindValue(t *testing.T) {

	tests := []struct {
		input    []string
		expected int64
	}{
		{
			[]string{"", "7,13,x,x,59,x,31,19"},
			1068781,
		}, {
			[]string{"", "67,7,59,61"},
			754018,
		}, {
			[]string{"", "67,x,7,59,61"},
			779210,
		}, {
			[]string{"", "67,7,x,59,61"},
			1261476,
		}, {
			[]string{"", "1789,37,47,1889"},
			1202161486,
		},
	}

	for _, test := range tests {
		t.Run(test.input[1], func(t *testing.T) {
			if val := findValue(test.input); val != test.expected {
				t.Fatalf("%d (%d)", val, test.expected)
			}

		})
	}
}
