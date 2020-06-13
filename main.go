package main

import (
	"app/game"

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

	gs := game.NewGameState(win)

	gs.AddScene("game-001", game.NewPIL001)
	gs.AddScene("game-002", game.NewPIL002)
	gs.AddScene("game-003", game.NewPIL003)

	gs.AddCharacter("hero", nil)

	id := "3"
	if len(os.Args) == 2 {
		id = os.Args[1]
	}

	if id == "1" {
		gs.ChangeScene("game-001")
	} else if id == "2" {
		gs.ChangeScene("game-002")
	} else if id == "3" {
		gs.ChangeScene("game-003")
	} else {
		panic("Unexpected game trying to load: " + id)
	}

	last := time.Now()
	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Greenyellow)

		gs.Update(dt)
		gs.Render(win)

		win.Update()
	}
}

func main() {
	pkger.Include("/characters.png")
	pkger.Include("/crater.png")
	pixelgl.Run(run)
}
