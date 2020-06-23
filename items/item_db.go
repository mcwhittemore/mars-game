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
	Sheet ItemSheet
	Icon  [2]float64
	Mind  MindFunc `json:"-"`
	Pos   pixel.Vec
}

func (item *Item) CanPickUp() bool {
	if item.Type == Plant_Type && item.Mind != nil {
		return false
	}
	return true
}

func GetSheet(id ItemSheet) *sheet.Sheet {
	return ItemSheets[id]
}

func (item *Item) GetSprite() (*pixel.Sprite, pixel.Matrix) {
	sheet := ItemSheets[item.Sheet]
	return sheet.GetSprite(item.Icon[0], item.Icon[1]), sheet.IM()
}

func (item *Item) PosBounds(pos pixel.Vec) pixel.Rect {
	sprite, im := item.GetSprite()
	bds := sprite.Frame()
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
	Seed_Type ItemType = iota
	Plant_Type
	Crop_Type
	Structure_Type
)

type ItemSheet int

const (
	Crop_Sheet ItemSheet = iota
	Wall_Sheet
	Conveyor_Sheet
)

var itemsDB = make(map[int]Item)
var itemIdxByName = make(map[string]int)

var cropSheet sheet.Sheet

var ItemSheets = make([]*sheet.Sheet, 0)

func init() {

	cropSheet, err := sheet.NewSheet("crops.png", pixel.Vec{X: 16, Y: 16}, pixel.ZV, sheet.TileSize/2)
	if err != nil {
		panic(err)
	}
	ItemSheets = append(ItemSheets, cropSheet)

	wallSheet, err := sheet.NewSheet("walls.png", pixel.Vec{X: 32, Y: 32}, pixel.ZV, sheet.TileSize)
	if err != nil {
		panic(err)
	}
	ItemSheets = append(ItemSheets, wallSheet)

	conveyorSheet, err := sheet.NewSheet("conveyorbelt.png", pixel.Vec{X: 32, Y: 32}, pixel.ZV, sheet.TileSize)
	if err != nil {
		panic(err)
	}
	ItemSheets = append(ItemSheets, conveyorSheet)

	itemsDB[0] = Item{
		Type:  Seed_Type,
		Name:  "Corn Seed",
		Class: "Corn",
		Sheet: Crop_Sheet,
		Icon:  [2]float64{5, 0},
	}

	itemsDB[1] = Item{
		Type:  Plant_Type,
		Name:  "Corn Plant",
		Class: "Corn",
		Sheet: Crop_Sheet,
		Icon:  [2]float64{4, 0},
	}

	itemsDB[2] = Item{
		Type:  Crop_Type,
		Name:  "Corn",
		Class: "Corn",
		Sheet: Crop_Sheet,
		Icon:  [2]float64{0, 0},
	}

	addItems(3, 3, 2, Wall_Sheet, Structure_Type, "Cinder Block", "Cinder Block %d")

	addItems(4, 1, 0, Conveyor_Sheet, Structure_Type, "Conveyor Belt", "Conveyor Belt %d")

	for i, item := range itemsDB {
		itemIdxByName[item.Name] = i
	}
}

func addItems(rows, cols, offset float64, sheet ItemSheet, t ItemType, class string, nameF string) {
	s := len(itemsDB)
	i := 0
	for x := float64(0); x < cols; x++ {
		for y := float64(0); y < rows; y++ {
			if y == 0 && x > cols-offset {
				continue
			}
			itemsDB[s+i] = Item{
				Type:  t,
				Name:  fmt.Sprintf(nameF, i),
				Class: class,
				Sheet: sheet,
				Icon:  [2]float64{x, y},
			}
			i++
		}
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
	if item.Type == Seed_Type {
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

	if item.Type == Plant_Type {
		return item.Class
	} else {
		return item.Name
	}
}
