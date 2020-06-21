package items

import (
	"app/sheet"

	"fmt"

	"github.com/faiface/pixel"
)

type Item struct {
	Type  ItemType
	Class string
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
	Structure
)

var itemsDB = make(map[int]Item)
var cropSheet sheet.Sheet
var itemIdxByName = make(map[string]int)

func init() {

	cropSheet, err := sheet.NewSheet("crops.png", pixel.Vec{X: 16, Y: 16}, pixel.ZV, sheet.TileSize/2)

	if err != nil {
		panic(err)
	}

	wallSheet, err := sheet.NewSheet("walls.png", pixel.Vec{X: 32, Y: 32}, pixel.ZV, sheet.TileSize)

	itemsDB[0] = Item{
		Type:  Seed,
		Name:  "Corn Seed",
		Class: "Corn",
		Sheet: cropSheet,
		Icon:  [2]float64{5, 0},
	}

	itemsDB[1] = Item{
		Type:  Plant,
		Name:  "Corn Plant",
		Class: "Corn",
		Sheet: cropSheet,
		Icon:  [2]float64{4, 0},
	}

	itemsDB[2] = Item{
		Type:  Crop,
		Name:  "Corn",
		Class: "Corn",
		Sheet: cropSheet,
		Icon:  [2]float64{0, 0},
	}

	itemsDB[3] = Item{
		Type:  Structure,
		Name:  "Cinder Block 001",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{0, 3},
	}

	itemsDB[4] = Item{
		Type:  Structure,
		Name:  "Cinder Block 002",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{1, 3},
	}

	itemsDB[5] = Item{
		Type:  Structure,
		Name:  "Cinder Block 003",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{2, 3},
	}

	itemsDB[6] = Item{
		Type:  Structure,
		Name:  "Cinder Block 004",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{0, 2},
	}

	itemsDB[7] = Item{
		Type:  Structure,
		Name:  "Cinder Block 005",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{1, 2},
	}

	itemsDB[8] = Item{
		Type:  Structure,
		Name:  "Cinder Block 006",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{2, 2},
	}

	itemsDB[9] = Item{
		Type:  Structure,
		Name:  "Cinder Block 006",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{0, 1},
	}

	itemsDB[10] = Item{
		Type:  Structure,
		Name:  "Cinder Block 008",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{1, 1},
	}

	itemsDB[11] = Item{
		Type:  Structure,
		Name:  "Cinder Block 009",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{2, 1},
	}

	itemsDB[12] = Item{
		Type:  Structure,
		Name:  "Cinder Block 010",
		Class: "Cinder Block",
		Sheet: wallSheet,
		Icon:  [2]float64{0, 0},
	}

	for i, item := range itemsDB {
		itemIdxByName[item.Name] = i
	}
}

func GetClassNames(class string) []string {
	names := make([]string, 0)

	for _, item := range itemsDB {
		if item.Class == class {
			names = append(names, item.Name)
		}
	}

	return names
}

func DropItem(name string, pos pixel.Vec) *Item {
	item := itemsDB[itemIdxByName[name]]
	if item.Type == Seed {
		item = itemsDB[itemIdxByName[fmt.Sprintf("%s Plant", item.Class)]]
		item.Mind = NewMindCropGrow()
	}
	item.Pos = pos
	return &item
}

func PickUpItem(item *Item) string {
	if item == nil || item.CanPickUp() == false {
		return ""
	}

	if item.Type == Plant {
		return item.Class
	} else {
		return item.Name
	}
}
