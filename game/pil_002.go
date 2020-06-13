package game

import (
	"app/characters"
	"app/maps"
	"app/sheet"

	"github.com/faiface/pixel"
)

type PIL002 struct {
	ground *maps.Map
	cam    pixel.Vec
}

func (g *PIL002) GetMap() *maps.Map {
	return g.ground
}

func (g *PIL002) GetCamera() pixel.Vec {
	return g.cam
}

func (g *PIL002) SetCamera(cam pixel.Vec) {
	g.cam = cam
}

func (g *PIL002) Enter(mi characters.MindInput) {
	mi.AddCharacter("alien", nil)
	mi.ShowCharacter("alien", characters.NewAlien(pixel.V(500, 500)))

	mi.AddCharacter("hero", nil)
	mi.ShowCharacter("hero", characters.NewHeroDefault(pixel.V(500, 600)))
}

func (g *PIL002) Exit(mi characters.MindInput) {
	mi.RemoveCharacter("alien")
	mi.RemoveCharacter("hero")
}

func (g *PIL002) Update(dt float64, mind characters.MindInput) {
	hp := mind.GetHeroPos()
	mind.KeepInView(hp, 200)
}

func NewPIL002() Scene {

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{X: 20, Y: 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

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
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 7, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 5, 0, 0, 0},
			{0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 9, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 2, 2, 2, 2, 2, 12, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 7, 6, 6, 6, 6, 6, 6, 6, 6, 5, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 10, 9, 0, 0, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 0, 0, 0, 10, 9, 0, 0, 4, 0, 11, 6, 6, 6, 6, 6, 5, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 10, 6, 6, 6, 12, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 4, 10, 9, 0, 0, 0, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 4, 11, 2, 2, 2, 5, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 11, 2, 2, 2, 2, 2, 2, 3, 0, 0, 0, 0, 4, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 9, 0, 4, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 4, 0, 4, 0, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 0, 4, 9, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 9, 0, 0, 0, 0, 11, 12, 0, 4, 0, 0, 0},
			{0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 4, 10, 9, 0, 0, 0, 0, 0, 4, 0, 0, 0},
			{0, 0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 1, 2, 2, 2, 2, 2, 2, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	})

	return &PIL002{
		ground: mapOne,
		cam:    pixel.V(600, 600),
	}

}
