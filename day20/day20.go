package main

import (
	"aoc2020/util"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type image struct {
	width  int
	height int
	bitmap []byte
}

func makeImage(width int, height int) image {
	img := image{width: width, height: height}
	img.bitmap = make([]byte, width*height)
	return img
}

func (img *image) ptr(x int, y int) *byte {
	return &img.bitmap[img.width*y+x]
}

func (img image) get(x int, y int) byte {
	return *img.ptr(x, y)
}

func (img *image) set(x int, y int, px byte) {
	*img.ptr(x, y) = px
}

func (img *image) blt(src image, dx int, dy int) {
	for y := 0; y < src.height; y++ {
		for x := 0; x < src.width; x++ {
			img.set(dx+x, dy+y, src.get(x, y))
		}
	}
}

func (img *image) bltMask(src image, dx int, dy int, c byte) {
	for y := 0; y < src.height; y++ {
		for x := 0; x < src.width; x++ {
			if src.get(x, y) == '#' {
				img.set(dx+x, dy+y, c)
			}
		}
	}
}

func (img *image) rotateCW() {
	newImg := makeImage(img.height, img.width)
	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			newImg.set(img.height-y-1, x, img.get(x, y))
		}
	}

	*img = newImg
}

func (img *image) flipVertical() {
	newImg := makeImage(img.width, img.height)
	for y := 0; y < img.height; y++ {
		for x := 0; x < img.width; x++ {
			newImg.set(x, y, img.get(x, img.height-y-1))
		}
	}
	*img = newImg
}

func (img image) countPixel(c byte) int {
	count := 0
	for _, px := range img.bitmap {
		if c == px {
			count++
		}
	}
	return count
}

func (img image) equal(b image) bool {
	if img.width == b.width {
		if img.height == b.height {
			return bytes.Equal(img.bitmap, b.bitmap)
		}
	}
	return false
}

const (
	upIndex = iota
	rightIndex
	downIndex
	leftIndex
)

type tileType struct {
	pixels      image
	adjacent    []int
	adjacentPtr [4]*tileType
}

type tileMap map[int]*tileType

func reverseString(input string) string {
	rev := ""
	for i := 0; i < len(input); i++ {
		rev += string(input[len(input)-i-1])
	}
	return rev
}

func parseBorder(input string) int {
	input = strings.ReplaceAll(input, "#", "1")
	input = strings.ReplaceAll(input, ".", "0")
	a, _ := strconv.ParseInt(input, 2, 32)
	b, _ := strconv.ParseInt(reverseString(input), 2, 32)
	if a < b {
		return int(a)
	}
	return int(b)
}

func parseTiles(input []string) tileMap {
	tiles := tileMap{}
	for i := 0; i < len(input); i++ {
		line := input[i]
		if line != "" {
			line = strings.ReplaceAll(line, "Tile ", "")
			line = strings.ReplaceAll(line, ":", "")
			val, _ := strconv.Atoi(line)

			i++

			lines := []string{}
			for j := 0; j < 10; j++ {
				lines = append(lines, input[i])
				i++
			}

			tile := tileType{}

			tile.adjacent = append(tile.adjacent, parseBorder(lines[0]))

			borderLeft := ""
			borderRight := ""
			for _, s := range lines {
				borderLeft += string(s[0])
				borderRight += string(s[9])
			}

			tile.pixels = makeImage(8, 8)
			for k := 0; k < 8; k++ {
				for l, c := range lines[k+1][1:9] {
					tile.pixels.set(l, k, byte(c))
				}
			}

			tile.adjacent = append(tile.adjacent, parseBorder(borderRight))
			tile.adjacent = append(tile.adjacent, parseBorder(lines[9]))
			tile.adjacent = append(tile.adjacent, parseBorder(borderLeft))

			tiles[val] = &tile
		}
	}

	tiles.connectTiles()
	return tiles
}

func findNeighbour(bordersA []int, bordersB []int) int {
	for i, a := range bordersA {
		for _, b := range bordersB {
			if a == b {
				return i
			}
		}
	}
	return -1
}

func (tiles *tileMap) connectTiles() {
	for outerTileID, outerTile := range *tiles {
		for innerTileID, innerTile := range *tiles {
			if outerTileID != innerTileID {
				if i := findNeighbour(outerTile.adjacent, innerTile.adjacent); i != -1 {
					outerTile.adjacentPtr[i] = innerTile
				}
			}
		}
	}
}

func (tiles tileMap) findCornerTiles() []int {
	cornerTiles := []int{}
	for tileID, val := range tiles {
		numAdjacent := 0
		for _, ptr := range val.adjacentPtr {
			if ptr != nil {
				numAdjacent++
			}
		}
		if numAdjacent == 2 {
			cornerTiles = append(cornerTiles, tileID)
		}
	}
	return cornerTiles
}

func part1(input []string) int {
	tiles := parseTiles(input)

	product := 1
	for _, id := range tiles.findCornerTiles() {
		product *= id
	}

	return product
}

func (tile *tileType) rotateCW() {
	tile.pixels.rotateCW()

	tile.adjacentPtr = [4]*tileType{
		tile.adjacentPtr[3],
		tile.adjacentPtr[0],
		tile.adjacentPtr[1],
		tile.adjacentPtr[2]}
}

func (tile *tileType) flipVertical() {
	tile.pixels.flipVertical()

	down := tile.adjacentPtr[downIndex]
	tile.adjacentPtr[downIndex] = tile.adjacentPtr[upIndex]
	tile.adjacentPtr[upIndex] = down
}

func (tile *tileType) adjustTile(left *tileType, up *tileType) {
	for tile.adjacentPtr[leftIndex] != left {
		tile.rotateCW()
	}

	if tile.adjacentPtr[upIndex] != up {
		tile.flipVertical()
	}
}

type onFindPattern func(img *image, pattern image, x int, y int)

func (img *image) findPattern(pattern image, onPattern onFindPattern) {
	for y := 0; y <= img.height-pattern.height; y++ {
		for x := 0; x <= img.width-pattern.width; x++ {

			scanPattern := func(x int, y int) bool {
				for py := 0; py < pattern.height; py++ {
					for px := 0; px < pattern.width; px++ {
						if pattern.get(px, py) == '#' {
							c := img.get(x+px, y+py)
							if c != '#' && c != 'O' {
								return false
							}
						}
					}
				}
				return true
			}

			if scanPattern(x, y) {
				onPattern(img, pattern, x, y)
			}
		}
	}
}

func makeImageFromText(input []string) image {
	img := makeImage(len(input[0]), len(input))
	for y, line := range input {
		for x, c := range line {
			img.set(x, y, byte(c))
		}
	}
	return img
}

func makeSeaMonster() image {
	seaMonsterInput := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}
	return makeImageFromText(seaMonsterInput)
}

func scanSeaForSeaMonsters(sea image) int {

	seaMonster := makeSeaMonster()
	scanSea := func(sea *image, seaMonster image) {
		sea.findPattern(seaMonster, func(img *image, pattern image, x int, y int) {
			img.bltMask(pattern, x, y, 'O')
		})
	}

	scanSeaRotate := func(sea *image, seaMonster *image) {
		scanSea(sea, *seaMonster)
		seaMonster.rotateCW()
		scanSea(sea, *seaMonster)
		seaMonster.rotateCW()
		scanSea(sea, *seaMonster)
		seaMonster.rotateCW()
		scanSea(sea, *seaMonster)
		seaMonster.rotateCW()
	}

	scanSeaRotate(&sea, &seaMonster)
	seaMonster.flipVertical()
	scanSeaRotate(&sea, &seaMonster)

	return sea.countPixel('#')
}

func part2(input []string) int {

	tiles := parseTiles(input)
	cornerTile := tiles.findCornerTiles()[0]
	tile, _ := tiles[cornerTile]

	var prevTile *tileType
	var upTile *tileType

	tilesOrdered := [][]*tileType{}

	tilesY := 0
	tilesX := 0
	for true {
		tilesX = 0

		tilesOrdered = append(tilesOrdered, []*tileType{})
		for true {
			if tilesY > 0 {
				upTile = tilesOrdered[tilesY-1][tilesX]
			} else {
				upTile = nil
			}

			tile.adjustTile(prevTile, upTile)
			tilesOrdered[tilesY] = append(tilesOrdered[tilesY], tile)
			prevTile = tile
			tile = tile.adjacentPtr[rightIndex]
			tilesX++
			if tile == nil {
				tile = tilesOrdered[tilesY][0].adjacentPtr[downIndex]
				prevTile = nil
				break
			}
		}

		tilesY++
		if tile == nil {
			break
		}
	}

	sea := makeImage(tilesX*8, tilesY*8)
	for y := 0; y < tilesY; y++ {
		for x := 0; x < tilesX; x++ {
			tile := tilesOrdered[y][x]
			sea.blt(tile.pixels, x*8, y*8)
		}
	}

	return scanSeaForSeaMonsters(sea)
}

func main() {
	input := util.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
