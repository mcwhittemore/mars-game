package game

import (
	"app/characters"
	"app/data"
	"app/items"
	"app/maps"

	"encoding/json"
	"io/ioutil"

	"github.com/faiface/pixel"
)

type PIL004 struct {
	ground *maps.Map
	cam    pixel.Vec
}

func (g *PIL004) GetMap() *maps.Map {
	return g.ground
}

func (g *PIL004) GetCamera() pixel.Vec {
	return g.cam
}

func (g *PIL004) SetCamera(cam pixel.Vec) {
	g.cam = cam
}

func (g *PIL004) Enter(mi characters.MindInput) {
	mi.ShowCharacter("hero", characters.NewHeroDefault(pixel.V(500, 600)))

	file, err := data.Open("/items/base-structure.json")
	items := make([]*items.Item, 0)

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &items)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		mi.AddItem(item)
	}
}

func (g *PIL004) Exit(mi characters.MindInput) {
	mi.HideCharacter("hero")
}

func (g *PIL004) Update(dt float64, mind characters.MindInput) {
	hp := mind.GetHeroPos()
	mind.KeepInView(hp, 200)
}

func NewPIL004() Scene {

	base := maps.NewMapFromFile("/maps/base.json")

	return &PIL004{
		ground: base,
		cam:    pixel.V(600, 600),
	}
}
