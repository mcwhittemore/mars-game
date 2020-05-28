package maps

import (
	"sheet"

	"github.com/faiface/pixel"
)

type MapOpts struct {
	Sheet *sheet.Sheet
	Tiles []*pixel.Vec
	Grid  [][]int
	Start pixel.Vec
}

func NewMap(opts *MapOpts) *pixel.Batch {

	dim := opts.Sheet.GetDim()

	sprites := make([]*pixel.Sprite, 0)

	for _, tile := range opts.Tiles {
		sprites = append(sprites, opts.Sheet.GetSprite(tile))
	}

	batch := opts.Sheet.GetBatch()

	right := pixel.Vec{dim, 0}

	for y, row := range opts.Grid {
		place := opts.Start.Add(pixel.Vec{0, float64(y) * dim})
		for _, tileId := range row {
			sprites[tileId].Draw(batch, opts.Sheet.IM().Moved(place))
			place = place.Add(right)
		}
	}

	return batch

}
