package characters

type CharacterData struct {
	Name      string
	Character *Character
}

func NewCharacterData(name string) *CharacterData {
	return &CharacterData{
		Name:      name,
		Character: nil,
	}
}
