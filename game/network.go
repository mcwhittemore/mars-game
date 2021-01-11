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

func inGrid(a Pos, cells [][]int) bool {
	if a.X < 0 || a.X >= len(cells) {
		return false
	}
	if a.Y < 0 || a.Y >= len(cells[0]) {
		return false
	}
	return true
}

func makeLine(p Line, e, n Pos, cells [][]int) (Line, bool, bool) {
	l := append(p, n)
	if inGrid(n, cells) && p.Contains(n) == false {
		if n.Eq(e) {
			return l, true, true
		} else {
			return l, true, false
		}
	} else {
		return l, false, false
	}
}

func addToList(q []Line, l Line) []Line {
	ll := make(Line, len(l))
	copy(ll, l)
	return append(q, ll)
}

func continueLine(p Line, e Pos, cells [][]int) ([]Line, bool, Line) {
	q := make([]Line, 0)
	done := false

	last := p[len(p)-1]

	win := make(Line, 0)

	locs := last.Adjacent()
	for _, np := range locs {
		nl, onGrid, winner := makeLine(p, e, np, cells)
		if onGrid {
			q = addToList(q, nl)
		}
		if winner {
			win = nl
			done = true
		}
	}

	return q, done, win
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func scoreLine(l Line, cells [][]int) int {
	return len(l)
	/*s := 0

	for _, p := range l {
		if cells[p.X][p.Y] != 1 {
			s += 10
		} else {
			s++
		}
	}

	return s*/
}

func splitOffShortest(q []Line, cells [][]int) ([]Line, Line) {
	idx := 0
	s := len(q[idx])

	l := len(q)
	for i := 1; i < l; i++ {
		is := scoreLine(q[i], cells)
		if s > is {
			idx = i
			s = is
		}
	}

	short := q[idx]

	q[idx] = q[l-1]
	return q[:l-1], short
}

func maybeAddToQueue(q []Line, l Line) bool {
	end := l[len(l)-1]
	for _, a := range q {
		if a.Contains(end) {
			return false
		}
	}
	return true
}

func (n *Network) shortestPath(a, b Pos, cells [][]int) Line {

	cells[b.X][b.Y]++

	p := make(Line, 1)
	p[0] = a
	q, done, win := continueLine(p, b, cells)

	var short Line
	var newLines []Line
	for done == false {
		q, short = splitOffShortest(q, cells)
		newLines, done, win = continueLine(short, b, cells)

		for _, nl := range newLines {
			if maybeAddToQueue(q, nl) {
				q = append(q, nl)
			}
		}

		if len(q) == 0 {
			return win
		}
	}

	return win
}

func findNearest(p Pos, cells [][]int) (Pos, bool) {
	near := Pos{
		X: len(cells),
		Y: len(cells[0]),
	}
	dist := len(cells) + len(cells[0])
	found := false
	for x := 0; x < len(cells); x++ {
		for y := 0; y < len(cells[x]); y++ {
			if cells[x][y] != 2 {
				continue
			}
			a := abs(x - p.X)
			b := abs(y - p.Y)
			if a == 0 && b == 0 {
				continue
			}
			d := int(math.Sqrt(float64(a*a + b*b)))
			if d < dist {
				found = true
				dist = d
				near.X = x
				near.Y = y
			}
		}
	}

	return near, found
}

func (n *Network) buildCells(builds []*Build, min pixel.Vec) [][]int {
	fmt.Println("Start buildCells")
	cells := make([][]int, int(n.area.W()))
	result := make([][]int, len(cells))
	for i := 0; i < int(n.area.W()); i++ {
		cells[i] = make([]int, int(n.area.H()))
		result[i] = make([]int, len(cells[i]))
	}

	l := len(builds)
	pos := make(Line, len(builds))
	for i, b := range builds {
		p := PosFromVec(b.Center().Sub(min))
		cells[p.X][p.Y] = 1
		pos[i] = p
	}

	if l > 1 {
		pts := n.shortestPath(pos[0], pos[1], cells)
		for _, p := range pts {
			result[p.X][p.Y]++
			cells[p.X][p.Y] = 2
		}
	}

	for i := 2; i < l; i++ {
		near, ok := findNearest(pos[i], cells)
		if ok {
			pts := n.shortestPath(pos[i], near, cells)
			for _, p := range pts {
				result[p.X][p.Y]++
				cells[p.X][p.Y] = 2
			}
		} else {
			panic("failed to find path target")
		}
	}

	fmt.Println("End buildCells")
	return result
}

func (n *Network) updatePaths(builds []*Build) {
	if n.area.Area() == 0 {
		n.grid = make([][]*Build, 0)
		return
	}

	cells := n.buildCells(builds, n.area.Min)

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
