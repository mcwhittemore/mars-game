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

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{20, 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	hero := NewHero()

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet: groundSheet,
		Tiles: []*pixel.Vec{{2, 6}, {0, 4}, {2, 4}, {4, 4}, {4, 6}, {4, 8}, {2, 8}, {0, 8}, {0, 6}},
		Grid: [][]int{
			{1, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{7, 6, 6, 6, 6, 6, 6, 6, 6, 5},
		},
		Start: pixel.ZV,
	})

	return &world.World{[]*pixel.Batch{mapOne}, []*characters.Character{}, hero, win.Bounds().Center(), win}
}
