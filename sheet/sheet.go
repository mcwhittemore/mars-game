package sheet

import (
	"github.com/faiface/pixel"
)

type Sheet struct {
	pic   pixel.Picture
	size  pixel.Vec
	base  pixel.Vec
	scale float64
	dim   float64
	cache map[float64]map[float64]*pixel.Sprite
}

func (s *Sheet) GetDim() float64 {
	return s.dim
}

func NewSheet(name string, s pixel.Vec, off pixel.Vec, dim float64) (*Sheet, error) {
	pic, err := loadPicture(name)
	if err != nil {
		return nil, err
	}

	bds := pic.Bounds()

	scale := dim / s.X
	vscale := dim / s.Y

	if vscale < scale {
		scale = vscale
	}

	xb := bds.Min.X + +off.X
	yb := bds.Min.Y + +off.Y

	base := pixel.Vec{X: xb, Y: yb}

	sheet := &Sheet{
		pic:   pic,
		size:  s,
		base:  base,
		scale: scale,
		dim:   dim,
	}

	sheet.cache = make(map[float64]map[float64]*pixel.Sprite, 0)

	return sheet, nil
}

func (s *Sheet) GetSprite(x, y float64) *pixel.Sprite {

	if _, ok := s.cache[x]; !ok {
		s.cache[x] = make(map[float64]*pixel.Sprite, 0)
	}

	if _, ok := s.cache[x][y]; !ok {
		xb := s.base.X + (s.size.X * x)
		yb := s.base.Y + (s.size.Y * y)

		sq := pixel.R(xb, yb, xb+s.size.X, yb+s.size.Y)

		s.cache[x][y] = pixel.NewSprite(s.pic, sq)
	}

	return s.cache[x][y]
}

func (s *Sheet) GetSprites(poss []pixel.Vec) []*pixel.Sprite {
	sprites := make([]*pixel.Sprite, 0)

	for _, pos := range poss {
		sprites = append(sprites, s.GetSprite(pos.X, pos.Y))
	}

	return sprites
}

func (s *Sheet) IM() pixel.Matrix {
	return pixel.IM.Scaled(pixel.ZV, s.scale)
}

func (s *Sheet) GetBatch() *pixel.Batch {
	return pixel.NewBatch(&pixel.TrianglesData{}, s.pic)
}
