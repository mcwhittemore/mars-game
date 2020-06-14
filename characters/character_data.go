package characters

import (
	"github.com/faiface/pixel/pixelgl"
)

type CharacterData struct {
	Name      string
	Character *Character
	items     map[string]int
	inHands   string
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
	cd.items[name]++
}

func NewCharacterData(name string) *CharacterData {
	items := make(map[string]int, 0)
	return &CharacterData{
		Name:      name,
		Character: nil,
		items:     items,
		inHands:   "",
	}
}
