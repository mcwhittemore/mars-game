package items

import (
	"app/sheet"

	"github.com/faiface/pixel"
)

type Item struct {
	Type  ItemType
	Name  string
	Sheet *sheet.Sheet
	Icon  [2]float64
	Mind  MindFunc
	Pos   pixel.Vec
}

func (item *Item) CanPickUp() bool {
	if item.Type == Plant && item.Mind != nil {
		return false
	}
	return true
}

func (item *Item) GetSprite() *pixel.Sprite {
	return item.Sheet.GetSprite(item.Icon[0], item.Icon[1])
}

func (item *Item) PosBounds(pos pixel.Vec) pixel.Rect {
	im := item.Sheet.IM()
	bds := item.GetSprite().Frame()
	bds.Min = im.Project(bds.Min)
	bds.Max = im.Project(bds.Max)
	w := (bds.Max.X - bds.Min.X) / 2
	h := (bds.Max.Y - bds.Min.Y) / 2

	return pixel.R(pos.X-w, pos.Y-h, pos.X+w, pos.Y+h)
}

func NewItem(name string, pos pixel.Vec, mind MindFunc) *Item {
	idx := itemIdxByName[name]
	item := itemsDB[idx]
	item.Pos = pos
	item.Mind = mind

	return &item
}

type ItemType int

const (
	Seed ItemType = iota
	Plant
	Crop
)

var itemsDB = make(map[int]Item)
var cropSheet sheet.Sheet
var itemIdxByName = make(map[string]int)

func init() {

	cropSheet, err := sheet.NewSheet("crops.png", pixel.Vec{X: 16, Y: 16}, pixel.Vec{X: 0, Y: 0}, 32)
	if err != nil {
		panic(err)
	}

	itemsDB[0] = Item{
		Type:  Seed,
		Name:  "Corn Seed",
		Sheet: cropSheet,
		Icon:  [2]float64{5, 0},
		Mind:  nil,
		Pos:   pixel.ZV,
	}

	itemsDB[1] = Item{
		Type:  Plant,
		Name:  "Corn Plant",
		Sheet: cropSheet,
		Icon:  [2]float64{4, 0},
		Mind:  nil,
		Pos:   pixel.ZV,
	}

	itemsDB[2] = Item{
		Type:  Crop,
		Name:  "Corn",
		Sheet: cropSheet,
		Icon:  [2]float64{0, 0},
		Mind:  nil,
		Pos:   pixel.ZV,
	}

	for i, item := range itemsDB {
		itemIdxByName[item.Name] = i
	}
}
