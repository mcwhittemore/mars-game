package game

import (
	"app/characters"
	"app/items"
	"app/maps"
	"app/sheet"

	"github.com/faiface/pixel"
)

type PIL003 struct {
	ground *maps.Map
	cam    pixel.Vec
}

func (g *PIL003) GetMap() *maps.Map {
	return g.ground
}

func (g *PIL003) GetCamera() pixel.Vec {
	return g.cam
}

func (g *PIL003) SetCamera(cam pixel.Vec) {
	g.cam = cam
}

func (g *PIL003) Enter(mi characters.MindInput) {
	mi.ShowCharacter("hero", characters.NewHeroDefault(pixel.V(500, 600)))

	mi.AddItem(items.NewItem("Corn Seed", pixel.V(600, 700), nil))
	mi.AddItem(items.NewItem("Corn Seed", pixel.V(600, 732), nil))
	mi.AddItem(items.NewItem("Corn Seed", pixel.V(600, 764), nil))
	mi.AddItem(items.NewItem("Corn Plant", pixel.V(664, 700), items.NewMindCropGrow()))
}

func (g *PIL003) Exit(mi characters.MindInput) {
	mi.HideCharacter("hero")
}

func (g *PIL003) Update(dt float64, mind characters.MindInput) {
	hp := mind.GetHeroPos()
	mind.KeepInView(hp, 200)
}

func NewPIL003() Scene {

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
			{7, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 5},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4},
			{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
		},
	})

	return &PIL003{
		ground: mapOne,
		cam:    pixel.V(600, 600),
	}
}
