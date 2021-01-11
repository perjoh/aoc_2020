package main

import (
	"aoc2020/util"
	"fmt"
)

type conwayCube [][][]byte

func makeConwayCube(sizeX int, sizeY int, sizeZ int) conwayCube {
	cube := make(conwayCube, sizeZ)
	for z := range cube {
		cube[z] = make([][]byte, sizeY)
		for y := range cube[z] {
			cube[z][y] = make([]byte, sizeX)
			for x := range cube[z][y] {
				cube[z][y][x] = '.'
			}
		}
	}
	return cube
}

func (cube conwayCube) size() (int, int, int) {
	return len(cube[0][0]), len(cube[0]), len(cube)

}

func (cube conwayCube) get(x int, y int, z int) byte {
	if 0 <= x && 0 <= y && 0 <= z {
		sx, sy, sz := cube.size()
		if x < sx && y < sy && z < sz {
			return cube[z][y][x]
		}
	}

	return '.'
}

func (cube *conwayCube) set(x int, y int, z int, c byte) {
	(*cube)[z][y][x] = c
}

func (cube conwayCube) determineState(cx int, cy int, cz int) byte {
	offsets := []int{-1, 0, 1}
	count := 0
	for _, z := range offsets {
		for _, y := range offsets {
			for _, x := range offsets {
				if cube.get(cx+x, cy+y, cz+z) == '#' {
					count++
				}
			}
		}
	}

	c := cube.get(cx, cy, cz)

	if c == '#' {
		count--
	}

	if (c == '.' && count == 3) || (c == '#' && (count == 2 || count == 3)) {
		return '#'
	}

	return '.'
}

func (cube conwayCube) runCycle(targetCube *conwayCube) {
	for z := range cube {
		for y := range cube[z] {
			for x := range cube[z][y] {
				state := cube.determineState(x, y, z)
				targetCube.set(x+1, y+1, z+1, state) // shift position to allow growth
			}
		}
	}
}

func (cube conwayCube) count(c byte) int {
	count := 0
	for z := range cube {
		for y := range cube[z] {
			for x := range cube[z][y] {
				if cube[z][y][x] == c {
					count++
				}
			}
		}
	}
	return count
}

func parseConwayCube(input []string) conwayCube {
	cube := makeConwayCube(len(input[0])+2, len(input)+2, 3)
	for y, line := range input {
		for x, c := range line {
			cube[1][y+1][x+1] = byte(c)
		}
	}
	return cube
}

func part1(input []string) int {
	cube := parseConwayCube(input)

	for cycle := 0; cycle < 6; cycle++ {
		cx, cy, cz := cube.size()
		cube2 := makeConwayCube(cx+2, cy+2, cz+2)
		cube.runCycle(&cube2)
		cube = cube2
	}

	return cube.count('#')
}

type hyperCube [][][][]byte

func makeHyperCube(sizeX int, sizeY int, sizeZ int, sizeW int) hyperCube {
	cube := make(hyperCube, sizeW)
	for w := range cube {
		cube[w] = make([][][]byte, sizeZ)
		for z := range cube[w] {
			cube[w][z] = make([][]byte, sizeY)
			for y := range cube[w][z] {
				cube[w][z][y] = make([]byte, sizeX)
				for x := range cube[w][z][y] {
					cube[w][z][y][x] = '.'
				}
			}
		}
	}
	return cube
}

func (cube hyperCube) sizeHC() (int, int, int, int) {
	return len(cube[0][0][0]), len(cube[0][0]), len(cube[0]), len(cube)
}

func (cube hyperCube) getHC(x int, y int, z int, w int) byte {
	if 0 <= x && 0 <= y && 0 <= z && 0 <= w {
		sx, sy, sz, sw := cube.sizeHC()
		if x < sx && y < sy && z < sz && w < sw {
			return cube[w][z][y][x]
		}
	}
	return '.'
}

func (cube *hyperCube) setHC(x int, y int, z int, w int, c byte) {
	(*cube)[w][z][y][x] = c
}

func (cube hyperCube) determineStateHC(cx int, cy int, cz int, cw int) byte {
	offsets := []int{-1, 0, 1}
	count := 0
	for _, w := range offsets {
		for _, z := range offsets {
			for _, y := range offsets {
				for _, x := range offsets {
					if cube.getHC(cx+x, cy+y, cz+z, cw+w) == '#' {
						count++
					}
				}
			}
		}
	}

	c := cube.getHC(cx, cy, cz, cw)

	if c == '#' {
		count--
	}

	if (c == '.' && count == 3) || (c == '#' && (count == 2 || count == 3)) {
		return '#'
	}

	return '.'
}

func (cube hyperCube) runCycleHC(targetCube *hyperCube) {
	for w := range cube {
		for z := range cube[w] {
			for y := range cube[w][z] {
				for x := range cube[w][z][y] {
					state := cube.determineStateHC(x, y, z, w)
					targetCube.setHC(x+1, y+1, z+1, w+1, state) // shift position to allow growth
				}
			}
		}
	}
}

func (cube hyperCube) countHC(c byte) int {
	count := 0
	for w := range cube {
		for z := range cube[w] {
			for y := range cube[w][z] {
				for x := range cube[w][z][y] {
					if cube[w][z][y][x] == c {
						count++
					}
				}
			}
		}
	}
	return count
}

func parseHyperCube(input []string) hyperCube {
	cube := makeHyperCube(len(input[0])+2, len(input)+2, 3, 3)
	for y, line := range input {
		for x, c := range line {
			cube[1][1][y+1][x+1] = byte(c)
		}
	}
	return cube
}

func part2(input []string) int {
	cube := parseHyperCube(input)

	for cycle := 0; cycle < 6; cycle++ {
		cx, cy, cz, cw := cube.sizeHC()
		cube2 := makeHyperCube(cx+2, cy+2, cz+2, cw+2)
		cube.runCycleHC(&cube2)
		cube = cube2
	}

	return cube.countHC('#')
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
