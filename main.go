package main

import (
	"app/world"

	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/markbates/pkger"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Mars Game!",
		Bounds: pixel.R(0, 0, 608, 608),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	mars := world.NewMars(win)

	last := time.Now()
	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Greenyellow)

		mars.Update(dt)
		mars.DrawHitBoxes()

		win.Update()
	}
}

func main() {
	pkger.Include("/characters.png")
	pkger.Include("/crater.png")
	pixelgl.Run(run)
}
