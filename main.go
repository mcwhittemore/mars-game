package main

import (
	"app/game"
	"app/perf"

	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	last := time.Now()
	perfOn := false
	if os.Getenv("PERFON") != "" {
		perfOn = true
	}

	g := game.NewGame(win)

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()
		g.Draw(dt)
		if perfOn {
			perf.Update()
		}
	}
}

func main() {
	pixelgl.Run(run)
}
