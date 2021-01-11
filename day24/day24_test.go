package main

import "testing"

func TestProcess(t *testing.T) {
	input := []string{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	}

	tiles := buildTiles(input)
	tiles = processTilesLong(tiles)
	if count := countTiles(tiles, false); count != 2208 {
		t.Fatalf("countTiles: %v(%v)", count, 2208)
	}
}
