package game

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Network struct {
	area  *Build
	paths []*Build
}

func NewNetwork() *Network {
	ac := pixel.RGB(0, 0, 1).Mul(pixel.Alpha(.3))
	area := NewBuild(pixel.ZV, pixel.V(0, 0), ac)
	n := Network{
		area: area,
	}
	return &n
}

func (n *Network) updateArea(builds []*Build) {
	bds := pixel.ZR
	if len(builds) > 0 {
		bds = builds[0].Bounds()
	} else {
		n.area.Move(bds.Min)
		n.area.Resize(bds.Max.Sub(bds.Min))
		return
	}

	for _, b := range builds {
		buildBds := b.Bounds()
		if bds.Min.X > buildBds.Min.X {
			bds.Min.X = buildBds.Min.X
		}
		if bds.Min.Y > buildBds.Min.Y {
			bds.Min.Y = buildBds.Min.Y
		}

		if bds.Max.X < buildBds.Max.X {
			bds.Max.X = buildBds.Max.X
		}
		if bds.Max.Y < buildBds.Max.Y {
			bds.Max.Y = buildBds.Max.Y
		}
	}

	bds.Min.X = bds.Min.X - 10
	bds.Min.Y = bds.Min.Y - 10
	bds.Max.X = bds.Max.X + 10
	bds.Max.Y = bds.Max.Y + 10

	n.area.Move(bds.Min)
	n.area.Resize(bds.Max.Sub(bds.Min))
}

func (n *Network) makePath(a, b *Build) []*Build {
	out := make([]*Build, 0)
	ac := a.Center()
	bc := b.Center()

	d := ac.Sub(bc).Floor()
	s := d.ScaledXY(d.Map(func(v float64) float64 {
		if v == 0 {
			return 0
		}
		return 1.0 / math.Abs(v)
	})).Floor()
	c := pixel.ZV
	clr := pixel.RGB(0, 0, 0)

	fmt.Printf("D: %v, S: %v\n", d, s)

	c.X += s.X
	for c.X != d.X {
		out = append(out, NewCell(ac.Sub(c).Floor(), clr))
		c.X += s.X
	}

	for c.Y != d.Y {
		out = append(out, NewCell(ac.Sub(c).Floor(), clr))
		c.Y += s.Y
	}

	return out
}

func (n *Network) updatePaths(builds []*Build) {
	n.paths = make([]*Build, 0)

	l := len(builds)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			n.paths = append(n.paths, n.makePath(builds[i], builds[j])...)
		}
	}
}

func (n *Network) Update(builds []*Build) {
	n.updateArea(builds)
	n.updatePaths(builds)
}

func (n *Network) Draw(w *pixelgl.Window) {
	n.area.Draw(w)
	for _, p := range n.paths {
		p.Draw(w)
	}
}
