package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	win     *pixelgl.Window
	imd     *imdraw.IMDraw
	mouse   *Build
	buildId int
	builds  []Build
}

func NewGame(w *pixelgl.Window) *Game {
	imd := imdraw.New(nil)
	mouse := NewBuild(w.MousePosition(), pixel.V(10, 10), pixel.RGB(1, 0, 0))
	g := Game{
		win:     w,
		imd:     imd,
		mouse:   mouse,
		buildId: 0,
		builds:  make([]Build, 0),
	}
	return &g
}

func (g *Game) Draw(dt float64) {
	g.win.Clear(colornames.Gray)
	g.imd.Clear()

	g.mouse.Move(g.win.MousePosition())
	g.mouse.Draw(g.win)
	for _, b := range g.builds {
		b.Draw(g.win)
	}

	g.win.Update()
}
