package main

import (
	"app/game001"
	"app/game002"
	"app/world"

	"os"
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

	id := "2"
	var game *world.World
	if len(os.Args) == 2 {
		id = os.Args[1]
	}

	if id == "1" {
		game = game001.NewMars(win)
	} else if id == "2" {
		game = game002.NewMars(win)
	} else {
		panic("Unexpected game trying to load: " + id)
	}

	last := time.Now()
	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Greenyellow)

		game.Update(dt)
		game.DrawHitBoxes()

		win.Update()
	}
}

func main() {
	pkger.Include("/characters.png")
	pkger.Include("/crater.png")
	pixelgl.Run(run)
}
