package game001

import (
	"app/characters"
	"app/game"
	"app/maps"
	"app/sheet"

	"github.com/faiface/pixel"
)

type Game001 struct {
	ground *maps.Map
	cam    pixel.Vec
}

func (g *Game001) GetMap() *maps.Map {
	return g.ground
}

func (g *Game001) GetCamera() pixel.Vec {
	return g.cam
}

func (g *Game001) SetCamera(cam pixel.Vec) {
	g.cam = cam
}

func (g *Game001) Enter(mi characters.MindInput) {
	safe := pixel.R(188, 200, 388, 400)

	ctr := safe.Center()

	mi.AddCharacter("rando-1", nil)
	mi.ShowCharacter("rando-1", NewRando(92, ctr))

	mi.AddCharacter("rando-2", nil)
	mi.ShowCharacter("rando-2", NewRando(89, ctr))

	mi.AddCharacter("rando-3", nil)
	mi.ShowCharacter("rando-3", NewRando(86, ctr))

	mi.AddCharacter("rando-4", nil)
	mi.ShowCharacter("rando-4", NewRando(83, ctr))

	mi.AddCharacter("hero", nil)
	mi.ShowCharacter("hero", NewHero())
}

func (g *Game001) Exit(mi characters.MindInput) {
	mi.RemoveCharacter("rando-1")
	mi.RemoveCharacter("rando-2")
	mi.RemoveCharacter("rando-3")
	mi.RemoveCharacter("rando-4")
}

func (g *Game001) Update(dt float64, mind characters.MindInput) {
}

func NewMars() game.Scene {

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{X: 20, Y: 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet:     groundSheet,
		Tiles:     []*maps.Tile{{2, 6}, {0, 4}, {2, 4}, {4, 4}, {4, 6}, {4, 8}, {2, 8}, {0, 8}, {0, 6}},
		TileTypes: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		Grid: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 7, 6, 6, 5, 0, 0, 0},
			{0, 0, 0, 8, 0, 0, 4, 0, 0, 0},
			{0, 0, 0, 8, 0, 0, 4, 0, 0, 0},
			{0, 0, 0, 1, 2, 2, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	})

	return &Game001{
		ground: mapOne,
		cam:    pixel.V(300, 300),
	}
}
