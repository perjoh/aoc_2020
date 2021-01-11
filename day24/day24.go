package main

import (
	"aoc2020/util"
	"fmt"
)

func step(dir string) (int, int) {
	/*e, +x
	  se, +x,-y
	  sw, -y
	  w, -x
	  nw, -x,+y
	  ne +y*/
	switch dir {
	case "e":
		return 1, 0
	case "se":
		return 1, -1
	case "sw":
		return 0, -1
	case "w":
		return -1, 0
	case "nw":
		return -1, 1
	case "ne":
		return 0, 1
	}
	return 0, 0
}

func toPoint(x, y int) int {
	return y*100000 + x
}

func traverse(path string) int {
	var x, y int

	for i := 0; i < len(path); i++ {
		dir := string(path[i])
		switch path[i] {
		case 's':
			dir += string(path[i+1])
			i++
		case 'n':
			dir += string(path[i+1])
			i++
		}

		dx, dy := step(dir)
		x += dx
		y += dy
	}

	return toPoint(x, y)
}

type tileMap map[int]bool

func buildTiles(input []string) tileMap {
	tiles := tileMap{}
	for _, path := range input {
		point := traverse(path)

		color, exist := tiles[point]
		if !exist {
			color = false
		} else {
			color = !color
		}
		tiles[point] = color
	}
	return tiles
}

func countTiles(tiles tileMap, color bool) int {
	count := 0
	for _, col := range tiles {
		if col == color {
			count++
		}
	}
	return count
}

func part1(input []string) int {
	tiles := buildTiles(input)
	return countTiles(tiles, false)
}

func stepPos(dir string) int {
	x, y := step(dir)
	return toPoint(x, y)
}

func getDirs() []string {
	return []string{"ne", "e", "se", "sw", "w", "nw"}
}

func countSurroundingTiles(tiles tileMap, pos int, color bool) int {
	dirs := getDirs()

	count := 0
	for _, dir := range dirs {
		col, exist := tiles[stepPos(dir)+pos]
		if !exist {
			col = true
		}

		if col == color {
			count++
		}
	}
	return count
}

func processTile(tilesIn tileMap, tilesOut *tileMap, pos int) {
	color, exist := tilesIn[pos]
	if !exist {
		color = true
	}

	(*tilesOut)[pos] = color

	if color == false {
		count := countSurroundingTiles(tilesIn, pos, false)
		if count == 0 || count > 2 {
			(*tilesOut)[pos] = true
		}
	} else {
		count := countSurroundingTiles(tilesIn, pos, false)
		if count == 2 {
			(*tilesOut)[pos] = false
		}
	}
}

func processAdjacentTiles(tilesIn tileMap, tilesOut *tileMap, pos int) {
	dirs := getDirs()
	for _, dir := range dirs {
		processTile(tilesIn, tilesOut, pos+stepPos(dir))
	}
}

func processTiles(tilesIn tileMap) tileMap {
	tilesOut := tileMap{}
	for pos, color := range tilesIn {
		processTile(tilesIn, &tilesOut, pos)
		if !color {
			processAdjacentTiles(tilesIn, &tilesOut, pos)
		}
	}
	return tilesOut
}

func processTilesLong(tiles tileMap) tileMap {
	for i := 0; i < 100; i++ {
		tiles = processTiles(tiles)
	}
	return tiles
}

func part2(input []string) int {
	tiles := buildTiles(input)
	tiles = processTilesLong(tiles)
	return countTiles(tiles, false)
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
