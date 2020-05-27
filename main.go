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
		if win.JustPressed(pixelgl.KeyD) {
			c.ChangePose("side")
		} else if win.JustPressed(pixelgl.KeyA) {
			c.ChangePose("left-side")
		} else if win.JustPressed(pixelgl.KeyS) {
			c.ChangePose("down")
		} else if win.JustPressed(pixelgl.KeyW) {
			c.ChangePose("up")
		}

		if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyW) {
			select {
			case <-second:
				c.Step()
			default:
			}
		} else {
			c.Stop()
		}

	})

	var offsetH, offsetV float64
	offsetH = 2 / 18
	offsetV = 2 / 20

	hero.AddPose("down", []pixel.Vec{{1, 95}, {2, 95}, {3 + offsetH, 95 - offsetV}, {4 + offsetH, 95}, {0, 95}}, pixel.Vec{0, -1})
	hero.AddPose("side", []pixel.Vec{{1, 96}, {2, 96}, {3 + offsetH, 96 - offsetV}, {4 + offsetH, 96}, {0, 96}}, pixel.Vec{1, 0})
	hero.AddPose("up", []pixel.Vec{{1, 97}, {2, 97}, {3 + offsetH, 97 - offsetV}, {4 + offsetH, 97}, {0, 97}}, pixel.Vec{0, 1})

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

		hs, hm := hero.Update(1)
		hs.Draw(win, hm)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
