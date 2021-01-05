package game

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	MOUSE_CLICK_TIME = 0.3
)

type Game struct {
	win           *pixelgl.Window
	camPos        pixel.Vec
	camZoom       float64
	mouseDownTime float64
	dragPos       pixel.Vec
	imd           *imdraw.IMDraw
	mouse         *Build
	network       *Network
	buildId       int
	builds        []*Build
}

func NewGame(w *pixelgl.Window) *Game {
	imd := imdraw.New(nil)
	mc := pixel.RGB(1, 0, 0).Mul(pixel.Alpha(.3))
	mouse := NewBuild(w.MousePosition(), pixel.V(1, 1), mc)

	network := NewNetwork()
	g := Game{
		win:           w,
		imd:           imd,
		mouseDownTime: 100.0,
		mouse:         mouse,
		network:       network,
		dragPos:       pixel.ZV,
		buildId:       0,
		builds:        make([]*Build, 0),
		camPos:        pixel.ZV,
		camZoom:       10.0,
	}
	return &g
}

func (g *Game) FindBuildByPos(p pixel.Vec) int {
	for i, b := range g.builds {
		if b.Contains(p) {
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

func (g *Game) AddBuild(p pixel.Vec) {
	nb := NewBuild(g.MousePosition(), pixel.V(1, 1), pixel.RGB(0, 1, 0))
	g.builds = append(g.builds, nb)
}

func (g *Game) Cam() pixel.Matrix {
	z := math.Floor(g.camZoom)
	return pixel.IM.Scaled(g.camPos, z).Moved(g.win.Bounds().Center().Sub(g.camPos))
}

func (g *Game) MousePosition() pixel.Vec {
	cam := g.Cam()
	return cam.Unproject(g.win.MousePosition()).Floor()
}

func (g *Game) MouseClicked() bool {
	if g.win.JustPressed(pixelgl.MouseButton1) {
		g.mouseDownTime = 0
	}

	return g.win.JustReleased(pixelgl.MouseButton1) && g.mouseDownTime < MOUSE_CLICK_TIME
}

func (g *Game) MouseDrag() pixel.Vec {
	if g.mouseDownTime > MOUSE_CLICK_TIME && g.win.Pressed(pixelgl.MouseButton1) {
		dp := g.win.MousePosition().Scaled(1 / g.camZoom)
		out := g.dragPos.Sub(dp)
		g.dragPos = dp
		return out
	} else {
		g.dragPos = g.win.MousePosition().Scaled(1 / g.camZoom)
		return pixel.ZV
	}
}

func (g *Game) Draw(dt float64) {
	g.mouseDownTime += dt
	md := g.MouseDrag()
	if md.Eq(pixel.ZV) == false {
		g.camPos = g.camPos.Add(md)
	}

	cam := g.Cam()
	g.win.SetMatrix(cam)
	g.win.Clear(colornames.Gray)
	g.imd.Clear()

	if g.MouseClicked() {
		b := g.FindBuildByPos(g.MousePosition())
		if b > -1 {
			g.RemoveBuild(b)
			g.network.Update(g.builds)
		} else {
			g.AddBuild(g.MousePosition())
			g.network.Update(g.builds)
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
	}

	for _, b := range g.builds {
		b.Draw(g.win)
	}
	g.mouse.Move(g.MousePosition())
	g.mouse.Draw(g.win)
	g.network.Draw(g.win)

	g.win.Update()
}
