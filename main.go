package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
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

	groundSheet, err := NewSheet("crater.png", pixel.Vec{20, 20}, pixel.ZV, 64)
	if err != nil {
		panic(err)
	}

	characterSheet, err := NewSheet("characters.png", pixel.Vec{18, 20}, pixel.Vec{0, 2}, 64)
	if err != nil {
		panic(err)
	}

	second := time.Tick(200 * time.Millisecond)
	hero := NewCharacter(characterSheet, win.Bounds().Center(), func(c *Character, dt int64) {
		select {
		case <-second:
			c.Step()
		default:
		}
	})

	hero.AddPose("down", []pixel.Vec{{0, 74}, {1, 74}, {2, 74}}, pixel.Vec{0, -1})

	hero.ChangePose("down")

	mapOne := NewMap(&MapOpts{
		Sheet: groundSheet,
		Tiles: []*pixel.Vec{{2, 6}},
		Grid:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		Start: win.Bounds().Center(),
	})

	for !win.Closed() {

		win.Clear(colornames.Greenyellow)

		mapOne.Draw(win)
		hs, hm := hero.Update(0)
		hs.Draw(win, hm)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
