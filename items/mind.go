package items

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type MindFunc func(*Item, float64, MindInput)

type Target interface {
	SetPos(pixel.Vec)
}

type MindInput interface {
	JustPressed(pixelgl.Button) bool
	Pressed(pixelgl.Button) bool
	//GetCollideRect(pixel.Rect, interface{}) (pixel.Rect, *Character)
	IsObstacle(pixel.Vec) bool
	GetHeroPos() pixel.Vec
	AddItem(*Item)
	GetLocation(string) pixel.Rect
	GetItems(pixel.Rect, func(*Item) pixel.Rect) []*Item
	GetTime() float64
	RemoveItem(*Item)
}
