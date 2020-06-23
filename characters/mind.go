package characters

import (
	"app/items"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

type MindInput interface {
	JustPressed(pixelgl.Button) bool
	Pressed(pixelgl.Button) bool
	Typed() string
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
	ListItems() []*items.Item
	AddDraw(*imdraw.IMDraw)
	GetLocation(string) pixel.Rect
	AddText(*text.Text)
}

type MindFunc func(*CharacterData, float64, MindInput)
