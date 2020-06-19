package maps

import (
	"app/data"
	"app/sheet"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/faiface/pixel"
)

type MapOpts struct {
	Sheet     *sheet.TileSheet
	Grid      [][]int
	Locations map[string]pixel.Rect
}

type Map struct {
	Render    *pixel.Batch
	Locations map[string]pixel.Rect
	tileDim   float64
	gridTypes [][]int
}

func (m *Map) IsObstacle(pos pixel.Vec) bool {
	x := int(pos.X / m.tileDim)
	y := int(pos.Y / m.tileDim)

	ly := len(m.gridTypes)
	lx := len(m.gridTypes[0])

	y = (ly - 1) - y

	if x >= 0 && x < lx && y >= 0 && y < ly {
		return m.gridTypes[y][x] != 0
	}
	return false
}

func NewMapFromFile(path string, sheet *sheet.TileSheet) *Map {
	opts := &MapOpts{}
	file, err := data.Open(path)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, opts)
	if err != nil {
		panic(err)
	}

	opts.Sheet = sheet
	return NewMap(opts)
}

func NewMap(opts *MapOpts) *Map {

	dim := opts.Sheet.Sheet.GetDim()

	sprites := make([]*pixel.Sprite, 0)

	ttlen := len(opts.Sheet.TileTypes)
	tlen := len(opts.Sheet.Tiles)

	if ttlen != tlen {
		panic(fmt.Sprintf("Expected Tiles (%d) and TileTypes (%d) to have the same length", tlen, ttlen))
	}

	for _, tile := range opts.Sheet.Tiles {
		sprites = append(sprites, opts.Sheet.Sheet.GetSprite(tile[0], tile[1]))
	}

	batch := opts.Sheet.Sheet.GetBatch()

	right := pixel.Vec{X: dim, Y: 0}

	gridTypes := make([][]int, len(opts.Grid))

	nr := len(opts.Grid) - 1
	for y, row := range opts.Grid {
		place := pixel.ZV.Add(pixel.Vec{X: 0, Y: float64(nr-y) * dim})
		rowTypes := make([]int, len(row))
		gridTypes[y] = rowTypes
		for x, tileId := range row {
			rowTypes[x] = opts.Sheet.TileTypes[tileId]
			sprites[tileId].Draw(batch, opts.Sheet.Sheet.IM().Moved(place))
			place = place.Add(right)
		}
	}

	loc := opts.Locations
	if loc == nil {
		loc = make(map[string]pixel.Rect)
	}

	return &Map{
		Render:    batch,
		tileDim:   dim,
		gridTypes: gridTypes,
		Locations: loc,
	}

}
