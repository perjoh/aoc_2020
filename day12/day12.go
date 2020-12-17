package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
)

type ship struct {
	pos struct {
		x int
		y int
	}

	// 0: east
	// 1: south
	// ...
	dir int
}

func toInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func (s *ship) rotate(instruction string) {
	steps := toInt(instruction[1:]) / 90
	if instruction[0] == 'L' {
		steps = -steps
	}
	s.dir = (4 + steps + s.dir) % 4
}

func (s *ship) runInstruction(instruction string) {
	// north, east: positive
	val := toInt(instruction[1:])
	switch instruction[0] {
	case 'N':
		s.pos.y += val
	case 'S':
		s.pos.y -= val
	case 'E':
		s.pos.x += val
	case 'W':
		s.pos.x -= val
	case 'L':
		s.rotate(instruction)
	case 'R':
		s.rotate(instruction)
	case 'F':
		switch s.dir {
		case 0:
			s.pos.x += val
		case 1:
			s.pos.y -= val
		case 2:
			s.pos.x -= val
		case 3:
			s.pos.y += val
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (s *ship) manhattanDistance() int {
	return abs(s.pos.x) + abs(s.pos.y)
}

func part1(input []string) int {
	s := ship{}
	for _, instruction := range input {
		s.runInstruction(instruction)
	}
	return s.manhattanDistance()
}

type shipExt struct {
	pos struct {
		x int
		y int
	}
	waypoint ship
}

func (s *shipExt) rotateWaypoint(instruction string) {
	steps := toInt(instruction[1:]) / 90
	if instruction[0] == 'R' {
		for steps > 0 {
			s.waypoint.pos.x, s.waypoint.pos.y = s.waypoint.pos.y, -s.waypoint.pos.x
			steps--
		}
	} else if instruction[0] == 'L' {
		for steps > 0 {
			s.waypoint.pos.x, s.waypoint.pos.y = -s.waypoint.pos.y, s.waypoint.pos.x
			steps--
		}
	}
}

func (s *shipExt) runInstruction(instruction string) {
	if instruction[0] == 'F' {
		steps := toInt(instruction[1:])
		s.pos.x += s.waypoint.pos.x * steps
		s.pos.y += s.waypoint.pos.y * steps
	} else if instruction[0] == 'R' || instruction[0] == 'L' {
		s.rotateWaypoint(instruction)
	} else {
		s.waypoint.runInstruction(instruction)
	}
}

func (s *shipExt) manhattanDistance() int {
	return abs(s.pos.x) + abs(s.pos.y)
}

func part2(input []string) int {
	s := shipExt{}
	s.waypoint.pos.x = 10
	s.waypoint.pos.y = 1
	for _, instr := range input {
		s.runInstruction(instr)
	}
	return s.manhattanDistance()
}

func main() {
	input := util.ReadStdin()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
