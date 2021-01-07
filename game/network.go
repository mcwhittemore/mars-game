package game

import (
	"fmt"

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

type Pos struct {
	X int
	Y int
}

func PosFromVec(p pixel.Vec) Pos {
	p = p.Floor()
	return Pos{
		X: int(p.X),
		Y: int(p.Y),
	}
}

func (p Pos) Eq(a Pos) bool {
	return p.X == a.X && p.Y == a.Y
}

func (p Pos) Up() Pos {
	return Pos{
		X: p.X,
		Y: p.Y - 1,
	}
}

func (p Pos) Down() Pos {
	return Pos{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p Pos) Left() Pos {
	return Pos{
		X: p.X - 1,
		Y: p.Y,
	}
}

func (p Pos) Right() Pos {
	return Pos{
		X: p.X + 1,
		Y: p.Y,
	}
}

func (p Pos) Sub(b Pos) Pos {
	return Pos{
		X: p.X - b.X,
		Y: p.Y - b.Y,
	}
}

func (n *Network) makePath(ac, bc Pos) []Pos {
	out := make([]Pos, 0)

	d := ac.Sub(bc)
	s := Pos{
		X: 1,
		Y: 1,
	}
	if d.X < 0 {
		s.X = -1
	}
	if d.Y < 0 {
		s.Y = -1
	}

	c := Pos{
		X: 0,
		Y: 0,
	}

	c.X += s.X
	for c.X != d.X {
		out = append(out, ac.Sub(c))
		c.X += s.X
	}

	for c.Y != d.Y {
		out = append(out, ac.Sub(c))
		c.Y += s.Y
	}

	return out
}

func isPath(a Pos, cells [][]int) bool {
	return cells[a.X][a.Y] > 0
}

func noLoop(a []Pos, l Pos) bool {
	for _, p := range a {
		if p.Eq(l) {
			return false
		}
	}
	return true
}

func checkAndAdd(q [][]Pos, p []Pos, e Pos, cells [][]int) ([][]Pos, bool, []Pos) {
	done := false

	last := p[len(p)-1]

	win := make([]Pos, 0)

	up := last.Up()
	if isPath(up, cells) && noLoop(p, up) {
		q = append(q, append(p, up))
		if up.Eq(e) {
			win = append(p, up)
			done = true
		}
	}

	down := last.Down()
	if isPath(down, cells) && noLoop(p, down) {
		q = append(q, append(p, down))
		if down.Eq(e) {
			win = append(p, down)
			done = true
		}
	}

	left := last.Left()
	if isPath(left, cells) && noLoop(p, left) {
		q = append(q, append(p, left))
		if left.Eq(e) {
			win = append(p, left)
			done = true
		}
	}

	right := last.Right()
	if isPath(right, cells) && noLoop(p, right) {
		q = append(q, append(p, right))
		if right.Eq(e) {
			win = append(p, right)
			done = true
		}
	}

	return q, done, win

}

func splitOffShortest(q [][]Pos) ([][]Pos, []Pos) {
	idx := 0
	s := len(q[idx])

	l := len(q)
	for i := 1; i < l; i++ {
		if s > len(q[i]) {
			idx = i
			s = len(q[i])
		}
	}

	short := q[idx]

	q[idx] = q[l-1]
	return q[:l-1], short
}

func findWinner(q [][]Pos, b Pos) []Pos {
	for _, p := range q {
		if p[len(p)-1].Eq(b) {
			return p
		}
	}
	panic("There should always be a winner")
}

func (n *Network) shortestPath(a, b Pos, cells [][]int) []Pos {

	cells[b.X][b.Y]++

	p := make([]Pos, 1)
	p[0] = a
	q, done, win := checkAndAdd(make([][]Pos, 0), p, b, cells)

	var short []Pos
	for done == false {
		q, short = splitOffShortest(q)
		q, done, win = checkAndAdd(q, short, b, cells)
		fmt.Printf("Q: %d\n", len(q))
		if len(q) == 0 {
			fmt.Printf("Fail: %v, %v\n", a, b)
			return n.makePath(a, b)
		}
	}

	if len(win) == 0 {
		panic("why no winner?")
	}

	return win
}

func (n *Network) buildCells(builds []*Build, min pixel.Vec) [][]int {
	cells := make([][]int, int(n.area.W()))
	result := make([][]int, len(cells))
	for i := 0; i < int(n.area.W()); i++ {
		cells[i] = make([]int, int(n.area.H()))
		result[i] = make([]int, len(cells[i]))
	}

	pos := make([]Pos, len(builds))
	for i, b := range builds {
		pos[i] = PosFromVec(b.Center().Sub(min))
	}

	l := len(builds)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			pts := n.makePath(pos[i], pos[j])
			for _, p := range pts {
				cells[p.X][p.Y]++
			}
		}
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			pts := n.shortestPath(pos[i], pos[j], cells)
			for _, p := range pts {
				result[p.X][p.Y]++
			}
		}
	}

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
