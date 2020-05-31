package world

import (
	"characters"
	"maps"
	"sheet"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewMars(win *pixelgl.Window) *World {

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{20, 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	safe := pixel.R(188, 200, 388, 400)
	npcs := make([]*characters.Character, 0)
	npcs = append(npcs, characters.NewRando(92, safe.Center()))
	npcs = append(npcs, characters.NewRando(89, safe.Center()))
	npcs = append(npcs, characters.NewRando(86, safe.Center()))
	npcs = append(npcs, characters.NewRando(83, safe.Center()))

	hero := characters.NewHero()

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet: groundSheet,
		Tiles: []*pixel.Vec{{2, 6}, {0, 4}, {2, 4}, {4, 4}, {4, 6}, {4, 8}, {2, 8}, {0, 8}, {0, 6}},
		Grid: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 2, 2, 3, 0, 0, 0},
			{0, 0, 0, 8, 0, 0, 4, 0, 0, 0},
			{0, 0, 0, 8, 0, 0, 4, 0, 0, 0},
			{0, 0, 0, 7, 6, 6, 5, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		Start: pixel.ZV,
	})

	return &World{[]*pixel.Batch{mapOne}, npcs, hero, win.Bounds().Center(), win}
}
