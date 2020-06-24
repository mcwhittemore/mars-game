package items

import (
	"app/sheet"

	"fmt"
	"strings"

	"github.com/faiface/pixel"
)

type Item struct {
	Type  ItemType
	Class string
	Name  string
	Sheet ItemSheet
	Icon  [2]float64
	State ItemState
	Pos   pixel.Vec
}

func (item *Item) CanPickUp() bool {
	if item.Type == Structure_Type {
		return false
	}

	if item.Type == Plant_Type && !item.State.UsingController("") {
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

func (item *Item) SetPos(pos pixel.Vec) {
	item.Pos = pos
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

func NewItem(name string, pos pixel.Vec, controller string) *Item {
	idx := itemIdxByName[name]
	item := itemsDB[idx]
	item.Pos = pos
	item.State = ItemState{
		Controller: controller,
		Data:       make(map[string]float64),
	}

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
	Spacesheet_Sheet
)

var itemsDB = make([]Item, 0)
var itemIdxByName = make(map[string]int)

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

	spaceshipSheet, err := sheet.NewSheet("spaceship.png", pixel.Vec{X: 96, Y: 96}, pixel.ZV, sheet.TileSize*9)
	if err != nil {
		panic(err)
	}
	ItemSheets = append(ItemSheets, spaceshipSheet)

	addCrop(0, "Corn")

	addItems(3, 3, 2, Wall_Sheet, Structure_Type, "Cinder Block", "Cinder Block %d")

	addItems(4, 1, 0, Conveyor_Sheet, Structure_Type, "Conveyor Belt", "Conveyor Belt %d")

	addItems(1, 1, 0, Spacesheet_Sheet, Structure_Type, "Spaceship", "Spaceship")

	for i, item := range itemsDB {
		itemIdxByName[item.Name] = i
	}
}

func addCrop(row float64, class string) {
	itemsDB = append(itemsDB, Item{
		Type:  Seed_Type,
		Name:  class + " Seed",
		Class: class,
		Sheet: Crop_Sheet,
		Icon:  [2]float64{5, row},
	})

	itemsDB = append(itemsDB, Item{
		Type:  Plant_Type,
		Name:  class + " Plant",
		Class: class,
		Sheet: Crop_Sheet,
		Icon:  [2]float64{4, row},
	})

	itemsDB = append(itemsDB, Item{
		Type:  Crop_Type,
		Name:  class,
		Class: class,
		Sheet: Crop_Sheet,
		Icon:  [2]float64{0, row},
	})
}

func addItems(rows, cols, offset float64, sheet ItemSheet, t ItemType, class string, nameF string) {
	i := 0
	for x := float64(0); x < cols; x++ {
		for y := float64(0); y < rows; y++ {
			if y == 0 && x > cols-offset {
				continue
			}

			name := nameF
			if strings.Count(nameF, "%d") > 0 {
				name = fmt.Sprintf(nameF, i)
			}

			itemsDB = append(itemsDB, Item{
				Type:  t,
				Name:  name,
				Class: class,
				Sheet: sheet,
				Icon:  [2]float64{x, y},
			})
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
		item.State.ChangeController("crop-grow")
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
