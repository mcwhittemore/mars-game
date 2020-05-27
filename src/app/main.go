package main

import (
	"characters"
	"sheet"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	groundSheet, err := sheet.NewSheet("crater.png", pixel.Vec{20, 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	hero := characters.NewHero(win)

	mapOne := NewMap(&MapOpts{
		Sheet: groundSheet,
		Tiles: []*pixel.Vec{{2, 6}},
		Grid:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		Start: win.Bounds().Center(),
	})

	for !win.Closed() {

		win.Clear(colornames.Greenyellow)

		mapOne.Draw(win)

		hs, hm := hero.Update(1)
		hs.Draw(win, hm)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
