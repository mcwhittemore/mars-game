package characters

import (
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
}

type MindFunc func(*Character, float64, MindInput)
