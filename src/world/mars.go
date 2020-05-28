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

	npcs := make([]*characters.Character, 0)
	npcs = append(npcs, characters.NewRando(92, pixel.Vec{32, 32}))
	npcs = append(npcs, characters.NewRando(89, pixel.Vec{-32, 32}))
	npcs = append(npcs, characters.NewRando(86, pixel.Vec{-32, -32}))
	npcs = append(npcs, characters.NewRando(83, pixel.Vec{32, -32}))

	hero := characters.NewHero()

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet: groundSheet,
		Tiles: []*pixel.Vec{{2, 6}},
		Grid:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		Start: pixel.ZV,
	})

	return &World{[]*pixel.Batch{mapOne}, npcs, hero, pixel.ZV, win}
}
