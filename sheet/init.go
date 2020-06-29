package sheet

import (
	"github.com/faiface/pixel"
)

var GroundTileSheet *TileSheet

const TileSize = float64(48)

func init() {

	groundSheet, err := NewSheet("crater.png", pixel.Vec{X: 20, Y: 20}, pixel.ZV, TileSize)
	if err != nil {
		panic(err)
	}

	GroundTileSheet = &TileSheet{
		Sheet:     groundSheet,
		Tiles:     []*Tile{{2, 6}, {0, 4}, {2, 4}, {4, 4}, {4, 6}, {4, 8}, {2, 8}, {0, 8}, {0, 6}, {2, 0}, {4, 0}, {4, 2}, {2, 2}},
		TileTypes: []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
}

func GetTileSheet(name string) *TileSheet {
	if name == "ground-tile-sheet" {
		return GroundTileSheet
	}
	panic("Unknown tilesheet: " + name)
}
