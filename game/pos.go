package game

import (
	"github.com/faiface/pixel"
)

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

func (p Pos) Adjacent() []Pos {
	out := make([]Pos, 4)
	out[0] = p.Up()
	out[1] = p.Right()
	out[2] = p.Down()
	out[3] = p.Left()
	return out
}

func (p Pos) Up() Pos {
	return Pos{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p Pos) Down() Pos {
	return Pos{
		X: p.X,
		Y: p.Y - 1,
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
