package main

import "testing"

func TestRunInstruction(t *testing.T) {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	s := ship{}
	for _, instruction := range instructions {
		s.runInstruction(instruction)
	}

	if md := s.manhattanDistance(); md != 25 {
		t.Fatalf("mc %d(%d)", md, 25)
	}
}

func TestWaypoint(t *testing.T) {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	s := shipExt{}
	s.waypoint.pos.x = 10
	s.waypoint.pos.y = 1
	for _, instr := range instructions {
		s.runInstruction(instr)
	}

	if md := s.manhattanDistance(); md != 286 {
		t.Fatalf("manhattanDistance: %d(%d)", md, 286)
	}

}
