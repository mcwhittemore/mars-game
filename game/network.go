package game

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var (
	CELL_COLORS []pixel.RGBA
)

func init() {
	CELL_COLORS = make([]pixel.RGBA, 3)
	CELL_COLORS[0] = pixel.RGB(0, 0, 1).Mul(pixel.Alpha(.3))
	CELL_COLORS[1] = pixel.RGB(0, 0, 0).Mul(pixel.Alpha(.6))
	CELL_COLORS[2] = pixel.RGB(0, 0, 0)
}

type Network struct {
	area pixel.Rect
	grid [][]*Build
}

func NewNetwork() *Network {
	area := pixel.ZR
	n := Network{
		area: area,
	}
	return &n
}

func (n *Network) updateArea(builds []*Build) {
	if len(builds) == 0 {
		n.area = pixel.ZR
		return
	}
	bds := builds[0].Bounds()

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

	bds.Min.X = bds.Min.X - 1
	bds.Min.Y = bds.Min.Y - 1
	bds.Max.X = bds.Max.X + 1
	bds.Max.Y = bds.Max.Y + 1

	n.area = bds
}

func (n *Network) makePath(a, b *Build, min pixel.Vec) []pixel.Vec {
	out := make([]pixel.Vec, 0)
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

	c.X += s.X
	for c.X != d.X {
		out = append(out, ac.Sub(c).Floor().Sub(min))
		c.X += s.X
	}

	for c.Y != d.Y {
		out = append(out, ac.Sub(c).Floor().Sub(min))
		c.Y += s.Y
	}

	return out
}

func (n *Network) buildCells(builds []*Build) [][]int {
	cells := make([][]int, int(n.area.W()))
	for i := 0; i < int(n.area.W()); i++ {
		cells[i] = make([]int, int(n.area.H()))
	}

	l := len(builds)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			pts := n.makePath(builds[i], builds[j], n.area.Min)
			for _, p := range pts {
				cells[int(p.X)][int(p.Y)]++
			}
		}
	}

	return cells
}

func (n *Network) updatePaths(builds []*Build) {
	if n.area.Area() == 0 {
		n.grid = make([][]*Build, 0)
		return
	}

	cells := n.buildCells(builds)

	h := 0
	w := len(n.grid)
	if w > 0 {
		h = len(n.grid[0])
	}

	flush := w != int(n.area.W()) || h != int(n.area.H())
	if flush {
		n.grid = make([][]*Build, int(n.area.W()))
	}
	fmt.Printf("Grid %dX%d was flushed? %v\n", int(n.area.W()), int(n.area.H()), flush)
	for x := 0; x < int(n.area.W()); x++ {
		if flush {
			n.grid[x] = make([]*Build, int(n.area.H()))
		}
		for y := 0; y < int(n.area.H()); y++ {
			c := cells[x][y]
			if c > 2 {
				c = 2
			}
			if flush {
				n.grid[x][y] = NewCell(pixel.V(float64(x), float64(y)).Add(n.area.Min), CELL_COLORS[c])
			} else {
				n.grid[x][y].Color(CELL_COLORS[c])
			}
		}
	}
}

func (n *Network) Update(builds []*Build) {
	n.updateArea(builds)
	n.updatePaths(builds)
}

func (n *Network) Draw(w *pixelgl.Window) {
	for _, r := range n.grid {
		for _, c := range r {
			c.Draw(w)
		}
	}
}
