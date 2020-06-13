package characters

type CharacterData struct {
	Name      string
	Character *Character
	items     map[int]int
}

func NewCharacterData(name string) *CharacterData {
	items := make(map[int]int, 0)
	return &CharacterData{
		Name:      name,
		Character: nil,
		items:     items,
	}
}
