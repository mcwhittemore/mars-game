package game

import (
	"app/characters"
	"app/maps"

	"github.com/faiface/pixel"
)

type PIL001 struct {
	ground *maps.Map
	cam    pixel.Vec
}

func (g *PIL001) GetMap() *maps.Map {
	return g.ground
}

func (g *PIL001) GetCamera() pixel.Vec {
	return g.cam
}

func (g *PIL001) SetCamera(cam pixel.Vec) {
	g.cam = cam
}

func (g *PIL001) Enter(mi characters.MindInput) {
	safe := pixel.R(188, 200, 388, 400)

	ctr := safe.Center()

	mi.AddCharacter("rando-1", nil)
	mi.ShowCharacter("rando-1", characters.NewRando(92, ctr))

	mi.AddCharacter("rando-2", nil)
	mi.ShowCharacter("rando-2", characters.NewRando(89, ctr))

	mi.AddCharacter("rando-3", nil)
	mi.ShowCharacter("rando-3", characters.NewRando(86, ctr))

	mi.AddCharacter("rando-4", nil)
	mi.ShowCharacter("rando-4", characters.NewRando(83, ctr))

	mi.ShowCharacter("hero", characters.NewHeroDefault(ctr))
}

func (g *PIL001) Exit(mi characters.MindInput) {
	mi.RemoveCharacter("rando-1")
	mi.RemoveCharacter("rando-2")
	mi.RemoveCharacter("rando-3")
	mi.RemoveCharacter("rando-4")
	mi.HideCharacter("rando-4")
}

func (g *PIL001) Update(dt float64, mi characters.MindInput) {
	safe := pixel.R(188, 200, 388, 400)
	hero := mi.GetCharacter("hero").Character
	isSafe := hero.Hits(safe)
	if isSafe {
		return
	}

	selfbox := hero.PosBounds(hero.Pos)
	_, subject := mi.GetCollideRect(selfbox, interface{}(hero))

	for subject != nil {
		subject.DropNear(safe.Center(), mi.GetCollideRect)
		_, subject = mi.GetCollideRect(selfbox, interface{}(hero))
	}
}

func NewPIL001() Scene {

	mapOne := maps.NewMap(&maps.MapOpts{
		Sheet: "ground-tile-sheet",
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

	return &PIL001{
		ground: mapOne,
		cam:    pixel.V(300, 300),
	}
}
