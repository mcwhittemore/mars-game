package main

import (
	"app/game"
	"app/perf"

	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	mon := pixelgl.PrimaryMonitor()
	w, h := mon.Size()
	cfg := pixelgl.WindowConfig{
		Title:     "Mars Base One",
		Bounds:    pixel.R(0, 0, w, h),
		Resizable: true,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	gs := game.NewGameState(win)

	gs.AddScene("game-001", game.NewPIL001)
	gs.AddScene("game-002", game.NewPIL002)
	gs.AddScene("game-003", game.NewPIL003)
	gs.AddScene("game-004", game.NewPIL004)
	gs.AddScene("world-builder", game.NewWorldBuilder)

	gs.AddCharacter("hero", nil)

	id := "4"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	if id == "world-builder" {
		gs.ChangeScene("world-builder")
	} else if id == "1" {
		gs.ChangeScene("game-001")
	} else if id == "2" {
		gs.ChangeScene("game-002")
	} else if id == "3" {
		gs.ChangeScene("game-003")
	} else if id == "4" {
		gs.ChangeScene("game-004")
	} else {
		panic("Unexpected game trying to load: " + id)
	}

	last := time.Now()

	perfOn := false
	if os.Getenv("PERFON") != "" {
		perfOn = true
	}

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Greenyellow)

		gs.Update(dt)
		gs.Render(win)

		win.Update()
		if perfOn {
			perf.Update()
		}
	}
}

func main() {
	pixelgl.Run(run)
}
