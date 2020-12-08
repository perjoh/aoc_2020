package main

import "testing"

func TestCalcSeatID(t *testing.T) {
	tests := []struct {
		bp  string
		row int
		col int
		sid int
	}{
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for _, test := range tests {
		t.Run(test.bp, func(t *testing.T) {
			sid := parseSeatID(test.bp)
			if sid != test.sid {
				t.Fatalf("%s, sid=%d(exp=%d)", test.bp, sid, test.sid)
			}
			if row := calcRow(sid); row != test.row {
				t.Fatalf("%s, row=%d(exp=%d)", test.bp, row, test.row)
			}
			if col := calcColumn(sid); col != test.col {
				t.Fatalf("%s, col=%d(exp=%d)", test.bp, col, test.col)
			}
		})
	}

}
