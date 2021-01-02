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
	builds  []*Build
}

func NewGame(w *pixelgl.Window) *Game {
	imd := imdraw.New(nil)
	mc := pixel.RGB(1, 0, 0).Mul(pixel.Alpha(.3))
	mouse := NewBuild(w.MousePosition(), pixel.V(10, 10), mc)
	g := Game{
		win:     w,
		imd:     imd,
		mouse:   mouse,
		buildId: 0,
		builds:  make([]*Build, 0),
	}
	return &g
}

func (g *Game) FindBuildByPos(p pixel.Vec) int {
	for i, b := range g.builds {
		if b.At(p) {
			return i
		}
	}
	return -1
}

func (g *Game) RemoveBuild(i int) {
	s := g.builds
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	g.builds = s[:len(s)-1]
}

func (g *Game) AddBuild(p pixel.Vec) bool {
	for _, b := range g.builds {
		if b.At(p) {
			return false
		}
	}
	nb := NewBuild(g.win.MousePosition(), pixel.V(10, 10), pixel.RGB(0, 1, 0))
	g.builds = append(g.builds, nb)
	return true
}

func (g *Game) Draw(dt float64) {
	g.win.Clear(colornames.Gray)
	g.imd.Clear()

	if g.win.JustPressed(pixelgl.MouseButton1) {
		b := g.FindBuildByPos(g.win.MousePosition())
		if b > -1 {
			g.RemoveBuild(b)
		} else {
			g.AddBuild(g.win.MousePosition())
		}
	}

	for _, b := range g.builds {
		b.Draw(g.win)
	}
	g.mouse.Move(g.win.MousePosition())
	g.mouse.Draw(g.win)

	g.win.Update()
}
