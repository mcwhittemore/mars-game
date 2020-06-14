package characters

import (
	"app/items"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type MindInput interface {
	JustPressed(pixelgl.Button) bool
	Pressed(pixelgl.Button) bool
	GetCollideRect(pixel.Rect, interface{}) (pixel.Rect, *Character)
	KeepInView(pixel.Vec, float64)
	IsObstacle(pixel.Vec) bool
	GetHeroPos() pixel.Vec
	AddCharacter(string, *CharacterData)
	ShowCharacter(string, *Character)
	RemoveCharacter(string)
	HideCharacter(string)
	GetCharacter(string) *CharacterData
	AddItem(*items.Item)
	GetItem(pixel.Vec) *items.Item
	RemoveItem(*items.Item)
}

type MindFunc func(*CharacterData, float64, MindInput)
