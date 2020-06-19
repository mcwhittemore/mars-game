package game

import (
	"app/characters"
	"app/maps"

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

	mi.ShowCharacter("hero", characters.NewHeroDefault(pixel.V(500, 600)))
}

func (g *PIL002) Exit(mi characters.MindInput) {
	mi.RemoveCharacter("alien")
	mi.HideCharacter("hero")
}

func (g *PIL002) Update(dt float64, mind characters.MindInput) {
	hp := mind.GetHeroPos()
	mind.KeepInView(hp, 200)
}

func NewPIL002() Scene {

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet: "ground-tile-sheet",
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
