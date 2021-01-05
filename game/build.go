package game

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Build struct {
	pos   pixel.Vec
	size  pixel.Vec
	color color.Color
	imd   *imdraw.IMDraw
}

func NewCell(p pixel.Vec, c color.Color) *Build {
	return NewBuild(p, pixel.V(1, 1), c)
}

func NewBuild(p, s pixel.Vec, c color.Color) *Build {
	imd := imdraw.New(nil)

	b := Build{
		size:  s,
		color: c,
		imd:   imd,
	}

	b.Move(p)
	b.rect()
	return &b
}

func (b *Build) Draw(w *pixelgl.Window) {
	b.imd.Draw(w)
}

func (b *Build) Bounds() pixel.Rect {
	return pixel.Rect{
		Min: b.pos,
		Max: b.pos.Add(b.size),
	}
}

func (b *Build) Center() pixel.Vec {
	return b.Bounds().Center()
}

func (b *Build) Contains(p pixel.Vec) bool {
	bds := b.Bounds()
	if bds.Area() == 1 {
		return p.Eq(b.pos)
	}
	return bds.Contains(p.Sub(pixel.V(.01, .01)))
}

func (b *Build) Resize(p pixel.Vec) {
	if p.Eq(b.size) == false {
		b.size = p
		b.rect()
	}
}

func (b *Build) Move(p pixel.Vec) {
	if p.Eq(b.pos) == false {
		b.pos = p
		b.rect()
	}
}

func (b *Build) rect() {
	b.imd.Clear()
	b.imd.Color = b.color
	b.imd.Push(b.pos)
	b.imd.Push(b.pos.Add(b.size))
	b.imd.Rectangle(0)
}
