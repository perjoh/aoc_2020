package main

import "testing"

func TestRotate(t *testing.T) {
	imgInput := []string{
		"####",
		"#...",
		".#..",
	}

	imgRotInput := []string{
		".##",
		"#.#",
		"..#",
		"..#",
	}

	img := makeImageFromText(imgInput)
	img.rotateCW()
	ok := img.equal(makeImageFromText(imgRotInput))
	if !ok {
		t.Fatalf("images not equal")
	}
}

func TestFlipVertical(t *testing.T) {
	imgInput := []string{
		"####",
		"#...",
		".#..",
	}

	imgRotInput := []string{
		".#..",
		"#...",
		"####",
	}

	img := makeImageFromText(imgInput)
	img.flipVertical()
	ok := img.equal(makeImageFromText(imgRotInput))
	if !ok {
		t.Fatalf("images not equal")
	}

}

func TestPixelCount(t *testing.T) {
	input := []string{
		".####...#####..#...###..",
		"#####..#..#.#.####..#.#.",
		".#.#...#.###...#.##.O#..",
		"#.O.##.OO#.#.OO.##.OOO##",
		"..#O.#O#.O##O..O.#O##.##",
		"...#.#..##.##...#..#..##",
		"#.##.#..#.#..#..##.#.#..",
		".###.##.....#...###.#...",
		"#.####.#.#....##.#..#.#.",
		"##...#..#....#..#...####",
		"..#.##...###..#.#####..#",
		"....#.##.#.#####....#...",
		"..##.##.###.....#.##..#.",
		"#...#...###..####....##.",
		".#.##...#.##.#.#.###...#",
		"#.###.#..####...##..#...",
		"#.###...#.##...#.##O###.",
		".O##.#OO.###OO##..OOO##.",
		"..O#.O..O..O.#O##O##.###",
		"#.#..##.########..#..##.",
		"#.#####..#.#...##..#....",
		"#....##..#.#########..##",
		"#...#.....#..##...###.##",
		"#..###....##.#...##.##.#",
	}

	sea := makeImageFromText(input)
	expected := 273
	if count := sea.countPixel('#'); count != expected {
		t.Fatalf("count %v(%v)", count, expected)
	}
}

func TestFindPattern(t *testing.T) {
	seaInput := []string{
		".####...#####..#...###..",
		"#####..#..#.#.####..#.#.",
		".#.#...#.###...#.##.##..",
		"#.#.##.###.#.##.##.#####",
		"..##.###.####..#.####.##",
		"...#.#..##.##...#..#..##",
		"#.##.#..#.#..#..##.#.#..",
		".###.##.....#...###.#...",
		"#.####.#.#....##.#..#.#.",
		"##...#..#....#..#...####",
		"..#.##...###..#.#####..#",
		"....#.##.#.#####....#...",
		"..##.##.###.....#.##..#.",
		"#...#...###..####....##.",
		".#.##...#.##.#.#.###...#",
		"#.###.#..####...##..#...",
		"#.###...#.##...#.######.",
		".###.###.#######..#####.",
		"..##.#..#..#.#######.###",
		"#.#..##.########..#..##.",
		"#.#####..#.#...##..#....",
		"#....##..#.#########..##",
		"#...#.....#..##...###.##",
		"#..###....##.#...##.##.#",
	}

	sea := makeImageFromText(seaInput)
	monster := makeSeaMonster()

	count := 0
	sea.findPattern(monster, func(img *image, pattern image, x, y int) {
		count++
	})

	if count != 2 {
		t.Fatalf("count=%v(%v)", count, 2)
	}
}

func TestRoughWaters(t *testing.T) {

	seaInput := []string{
		".#.#..#.##...#.##..#####",
		"###....#.#....#..#......",
		"##.##.###.#.#..######...",
		"###.#####...#.#####.#..#",
		"##.#....#.##.####...#.##",
		"...########.#....#####.#",
		"....#..#...##..#.#.###..",
		".####...#..#.....#......",
		"#..#.##..#..###.#.##....",
		"#.####..#.####.#.#.###..",
		"###.#.#...#.######.#..##",
		"#.####....##..########.#",
		"##..##.#...#...#.#.#.#..",
		"...#..#..#.#.##..###.###",
		".#.#....#.##.#...###.##.",
		"###.#...#..#.##.######..",
		".#.#.###.##.##.#..#.##..",
		".####.###.#...###.#..#.#",
		"..#.#..#..#.#.#.####.###",
		"#..####...#.#.#.###.###.",
		"#####..#####...###....##",
		"#.##..#..#...#..####...#",
		".#.###..##..##..####.##.",
		"...###...##...#...#..###",
	}

	sea := makeImageFromText(seaInput)

	roughness := scanSeaForSeaMonsters(sea)
	if roughness != 273 {
		t.Fatalf("roughness=%v (%v)", roughness, 273)
	}
}