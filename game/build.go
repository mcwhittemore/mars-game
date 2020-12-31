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

func NewBuild(p, s pixel.Vec, c color.Color) *Build {
	imd := imdraw.New(nil)

	b := Build{
		pos:   p,
		size:  s,
		color: c,
		imd:   imd,
	}

	b.rect()
	return &b
}

func (b *Build) Draw(w *pixelgl.Window) {
	b.imd.Draw(w)
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
