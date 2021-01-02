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

func (b *Build) At(p pixel.Vec) bool {
	m := 10.0
	v := p.Scaled(1.0 / m).Floor().Scaled(m)
	return v.Eq(b.pos)
}

func (b *Build) Move(p pixel.Vec) {
	m := 10.0
	v := p.Scaled(1.0 / m).Floor().Scaled(m)
	if v.Eq(b.pos) == false {
		b.pos = v
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
