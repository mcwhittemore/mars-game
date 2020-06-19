package sheet

type Tile [2]float64

type TileSheet struct {
	Sheet     *Sheet
	Tiles     []*Tile
	TileTypes []int
}
