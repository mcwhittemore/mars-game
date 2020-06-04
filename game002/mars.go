package game002

import (
	"app/characters"
	"app/maps"
	"app/sheet"
	"app/world"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewMars(win *pixelgl.Window) *world.World {

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{X: 20, Y: 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	hero := NewHero()

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet:     groundSheet,
		Tiles:     []*maps.Tile{{2, 6}, {0, 4}, {2, 4}, {4, 4}, {4, 6}, {4, 8}, {2, 8}, {0, 8}, {0, 6}, {2, 0}, {4, 0}, {4, 2}, {2, 2}},
		TileTypes: []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		// 0: empty
		// 1: |_
		// 2: __
		// 3: _|
		// 4:  |
		// 5: ⎻|
		// 6: ⎻⎻
		// 7: |⎻
		// 8: |
		// 9:  ⎻| alt
		// 10: |⎻ alt
		// 11: |_ alt
		// 12: _| alt
		Grid: [][]int{
			{7, 6, 6, 6, 6, 6, 6, 6, 6, 5},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 10, 9, 0, 0, 4},
			{8, 0, 10, 6, 6, 6, 12, 0, 0, 4},
			{8, 0, 8, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 8, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 11, 2, 2, 2, 2, 2, 2, 3},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
		Start: pixel.ZV,
	})

	return &world.World{[]*maps.Map{mapOne}, []*characters.Character{}, hero, win.Bounds().Center(), win}
}
