package maps

import (
	"app/sheet"

	"github.com/faiface/pixel"
)

type Tile [3]float64

type MapOpts struct {
	Sheet *sheet.Sheet
	Tiles []*Tile
	Grid  [][]int
	Start pixel.Vec
}

func NewMap(opts *MapOpts) *pixel.Batch {

	dim := opts.Sheet.GetDim()

	sprites := make([]*pixel.Sprite, 0)

	for _, tile := range opts.Tiles {
		sprites = append(sprites, opts.Sheet.GetSprite(tile[0], tile[1]))
	}

	batch := opts.Sheet.GetBatch()

	right := pixel.Vec{X: dim, Y: 0}

	nr := len(opts.Grid) - 1
	for y, row := range opts.Grid {
		place := opts.Start.Add(pixel.Vec{X: 0, Y: float64(nr-y) * dim})
		for _, tileId := range row {
			sprites[tileId].Draw(batch, opts.Sheet.IM().Moved(place))
			place = place.Add(right)
		}
	}

	return batch

}
