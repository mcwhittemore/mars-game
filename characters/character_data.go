package characters

import (
	"fmt"
	"math/rand"

	"github.com/faiface/pixel/pixelgl"
)

type CharacterItem struct {
	Name  string
	Count int
}

type CharacterData struct {
	Name      string
	Character *Character
	ItemsEtag string
	Items     []CharacterItem
	InHands   int
}

func (cd *CharacterData) Update(dt float64, mi MindInput) {
	if cd.Character != nil {
		cd.Character.mind(cd, dt, mi)
	}
}

func (cd *CharacterData) Render(win *pixelgl.Window) {
	if cd.Character != nil {
		cd.Character.Render(win)
	}
}

func (cd *CharacterData) AddItem(name string) {
	cd.ItemsEtag = fmt.Sprintf("%d", rand.Int())
	for i, ci := range cd.Items {
		if ci.Name == name {
			cd.Items[i].Count++
			return
		}
	}

	cd.InHands = 0
	ci := CharacterItem{
		Name:  name,
		Count: 1,
	}

	cd.Items = append(cd.Items, ci)
}

func (cd *CharacterData) RemoveItem(name string) {
	cd.ItemsEtag = fmt.Sprintf("%d", rand.Int())
	for i, ci := range cd.Items {
		if ci.Name == name {
			cd.Items[i].Count--

			if ci.Count == 0 {
				cd.Items = append(cd.Items[:i], cd.Items[i+1:]...)

				ni := len(cd.Items)

				if ni > i {
					cd.InHands = i
				} else {
					cd.InHands = i - 1
				}
			}

			return
		}
	}
}

func NewCharacterData(name string) *CharacterData {
	items := make([]CharacterItem, 0)
	return &CharacterData{
		Name:      name,
		Character: nil,
		Items:     items,
		InHands:   -1,
	}
}
