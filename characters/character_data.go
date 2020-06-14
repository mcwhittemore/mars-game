package characters

import (
	"github.com/faiface/pixel/pixelgl"
)

type CharacterData struct {
	Name      string
	Character *Character
	items     map[string]int
	InHands   string
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
	cd.items[name] = cd.items[name] + 1
	if cd.InHands == "" {
		cd.InHands = name
	}
}

func (cd *CharacterData) RemoveItem(name string) {
	cd.items[name]--
	if cd.items[name] == 0 {
		delete(cd.items, name)
		cd.InHands = ""
		for k := range cd.items {
			cd.InHands = k
			return
		}
	}
}

func NewCharacterData(name string) *CharacterData {
	items := make(map[string]int, 0)
	return &CharacterData{
		Name:      name,
		Character: nil,
		items:     items,
		InHands:   "",
	}
}
