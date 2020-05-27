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

	beings := make([]*characters.Character, 0)
	beings = append(beings, characters.NewHero(win))
	beings = append(beings, characters.NewRando(win, 92, win.Bounds().Center().Add(pixel.Vec{32, 32})))
	beings = append(beings, characters.NewRando(win, 89, win.Bounds().Center().Add(pixel.Vec{-32, 32})))
	beings = append(beings, characters.NewRando(win, 86, win.Bounds().Center().Add(pixel.Vec{-32, -32})))
	beings = append(beings, characters.NewRando(win, 83, win.Bounds().Center().Add(pixel.Vec{32, -32})))

	mapOne := NewMap(&MapOpts{
		Sheet: groundSheet,
		Tiles: []*pixel.Vec{{2, 6}},
		Grid:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		Start: win.Bounds().Center(),
	})

	for !win.Closed() {

		win.Clear(colornames.Greenyellow)

		mapOne.Draw(win)

		for _, being := range beings {
			bs, bm := being.Update(1)
			bs.Draw(win, bm)
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
