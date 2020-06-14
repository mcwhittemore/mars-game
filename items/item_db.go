package items

import (
	"app/characters"
	"app/sheet"

	"github.com/faiface/pixel"
)

type Item struct {
	Id    int
	Type  ItemType
	Name  string
	Sheet *sheet.Sheet
	Icon  [2]int
	Life  func(*Item, *characters.MindInput)
}

type ItemType int

const (
	Seed ItemType = iota
	Plant
	Crop
)

var ItemsDB = make(map[int]Item)
var cropSheet sheet.Sheet

func init() {

	cropSheet, err := sheet.NewSheet("crops.png", pixel.Vec{X: 16, Y: 16}, pixel.Vec{X: 0, Y: 0}, 64)
	if err != nil {
		panic(err)
	}

	ItemsDB[0] = Item{
		Id:    0,
		Type:  Seed,
		Name:  "Corn Seed",
		Sheet: cropSheet,
		Icon:  [2]int{5, 0},
		Life:  nil,
	}

	ItemsDB[1] = Item{
		Id:    1,
		Type:  Plant,
		Name:  "Corn Plant",
		Sheet: cropSheet,
		Icon:  [2]int{4, 0},
		Life:  nil,
	}

	ItemsDB[2] = Item{
		Id:    2,
		Type:  Crop,
		Name:  "Corn",
		Sheet: cropSheet,
		Icon:  [2]int{0, 0},
		Life:  nil,
	}
}
