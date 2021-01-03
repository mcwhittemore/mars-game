package game

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	win     *pixelgl.Window
	camPos  pixel.Vec
	camZoom float64
	imd     *imdraw.IMDraw
	mouse   *Build
	buildId int
	builds  []*Build
}

func NewGame(w *pixelgl.Window) *Game {
	imd := imdraw.New(nil)
	mc := pixel.RGB(1, 0, 0).Mul(pixel.Alpha(.3))
	mouse := NewBuild(w.MousePosition(), pixel.V(1, 1), mc)
	g := Game{
		win:     w,
		imd:     imd,
		mouse:   mouse,
		buildId: 0,
		builds:  make([]*Build, 0),
		camPos:  pixel.ZV,
		camZoom: 10.0,
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
	nb := NewBuild(g.MousePosition(), pixel.V(1, 1), pixel.RGB(0, 1, 0))
	g.builds = append(g.builds, nb)
	return true
}

func (g *Game) Cam() pixel.Matrix {
	z := math.Floor(g.camZoom)
	return pixel.IM.Scaled(g.camPos, z).Moved(g.win.Bounds().Center().Sub(g.camPos))
}

func (g *Game) MousePosition() pixel.Vec {
	cam := g.Cam()
	return cam.Unproject(g.win.MousePosition()).Floor()
}

func (g *Game) Draw(dt float64) {
	cam := g.Cam()
	g.win.SetMatrix(cam)
	g.win.Clear(colornames.Gray)
	g.imd.Clear()

	if g.win.JustPressed(pixelgl.MouseButton1) {
		b := g.FindBuildByPos(g.MousePosition())
		if b > -1 {
			g.RemoveBuild(b)
		} else {
			g.AddBuild(g.MousePosition())
		}
	}

	ms := g.win.MouseScroll()
	if pixel.ZV.Eq(ms) == false {
		g.camZoom += ms.Y
		if g.camZoom < 1 {
			g.camZoom = 1
		} else if g.camZoom > 20 {
			g.camZoom = 20
		}
		fmt.Printf("Scroll: %f, %f\n", ms.Y, g.camZoom)
	}

	for _, b := range g.builds {
		b.Draw(g.win)
	}
	g.mouse.Move(g.MousePosition())
	g.mouse.Draw(g.win)

	g.win.Update()
}
